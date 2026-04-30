package memory

import (
	"syscall"
	"unsafe"
)

// the structure
// header of size 8 bytes | user data | footer of size 8 bytes

type ttags struct {
	ptr  uintptr
	size uint32
}

var N int32 = 100 * 1024 * 1024 //first allocation of the heap in bytes
var heap []byte                 //heap memory

var freememory = make(map[ttags]bool) //first element points to the starting index of the block and

var heapstrtindex uintptr = 0
var heapendindex uintptr = 0
var currheapidx uintptr = 0

func initheap(size int) {
	heap, err := syscall.Mmap(-1, 0, size, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_PRIVATE|syscall.MAP_ANON)

	if err != nil {
		panic(err)
	}
	heapstrtindex = uintptr(unsafe.Pointer(&heap[0]))
	heapendindex = heapstrtindex + uintptr(size)
	currheapidx = heapstrtindex
}

// write the tag of the block
// @param(sz is the size of the block)
// @param(isfree is true if the block is free)
func write_tag(sz uint32, addr uintptr, isfree bool) {
	val := sz
	if isfree {
		val |= 1
	} else {
		val &= ^uint32(1)
	}
	// binary.LittleEndian.PutUint32(heap[cur_byte_idx:cur_byte_idx+4], val)

	// cur_byte_idx += 4
	curr := unsafe.Pointer(addr)
	if uintptr(curr)+4 > heapendindex {
		panic("out of memory")
	}
	*(*uint32)(curr) = val
}

// for now we will do a linear search to find the memory(req)
func find_mem(size uint32) uintptr {
	for x := range freememory {
		if x.size >= size {
			return x.ptr
		}
	}
	return 0 // nil
}

func Malloc(size uint) uintptr {
	//first check if there is any free block of required size
	//but before that we will round of this size to the next power of 8
	if (size % 8) != 0 {
		var rem = size / 8
		size = (rem + 1) * 8
	}

	size += 8

	//we will traverse the freememory map and check if there is any block of required size

	//first check if a pointer is available or not
	memptr := find_mem(uint32(size))

	if memptr != 0 {
		return memptr
	}

	//check if memeory is left
	if currheapidx+uintptr(size) > heapendindex {
		panic("out of memory")
	}

	//else allocate the required memeory
	ansptr := currheapidx + 4

	//header
	write_tag(uint32(size), currheapidx, false)

	currheapidx += uintptr(size)

	//footer
	write_tag(uint32(size), currheapidx-4, false)

	return ansptr

}

func Free(ptr uintptr) {
	//what is the size of the block
	szptr := ptr - 4 //header is of size 4 bytes
	sz := *(*uint32)(unsafe.Pointer(szptr))
	sz = sz & ^uint32(1) //extracting the size
	lo := szptr
	hi := ptr + uintptr(sz)

	//first we have to check if we can join the size ways of the memory
	//left
	footerleftsz := lo - 4

	if uintptr(footerleftsz) >= uintptr(unsafe.Pointer(&heap[0])) && *(*uint32)(unsafe.Pointer(footerleftsz))&1 == 0 {
		leftsz := *(*uint32)(unsafe.Pointer(footerleftsz))
		leftsz = leftsz & ^uint32(1)
		sz += leftsz

		szptr -= uintptr(leftsz)
		write_tag(uint32(sz), uintptr(szptr), true)
		write_tag(uint32(sz), uintptr(szptr)+uintptr(sz)-4, true)
	}
	//right
	headerrightsz := hi
	if uintptr(headerrightsz) < heapendindex && *(*uint32)(unsafe.Pointer(headerrightsz))&1 == 0 {
		rightsz := *(*uint32)(unsafe.Pointer(headerrightsz))
		rightsz = rightsz & ^uint32(1)
		sz += rightsz

		write_tag(uint32(sz), uintptr(szptr), true)
		write_tag(uint32(sz), uintptr(szptr)+uintptr(sz)-4, true)
	}
}
