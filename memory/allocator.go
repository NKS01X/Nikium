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
	var err error
	heap, err = syscall.Mmap(-1, 0, size, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_PRIVATE|syscall.MAP_ANON)

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
func find_mem(size uint32) ttags {
	for x := range freememory {
		if x.size >= size {
			return x
		}
	}
	return ttags{}
}

func Malloc(size uint) uintptr {
	if (size % 8) != 0 {
		var rem = size / 8
		size = (rem + 1) * 8
	}

	size += 8

	tag := find_mem(uint32(size))

	if tag.ptr != 0 {
		delete(freememory, tag)
		// Splitting
		if tag.size >= uint32(size)+16 {
			new_sz := tag.size - uint32(size)
			rem_ptr := tag.ptr + uintptr(size)

			write_tag(uint32(size), tag.ptr-4, false)
			write_tag(uint32(size), tag.ptr-4+uintptr(size)-4, false)

			write_tag(new_sz, rem_ptr-4, true)
			write_tag(new_sz, rem_ptr-4+uintptr(new_sz)-4, true)
			freememory[ttags{ptr: rem_ptr, size: new_sz}] = true
			return tag.ptr
		}
		write_tag(tag.size, tag.ptr-4, false)
		write_tag(tag.size, tag.ptr-4+uintptr(tag.size)-4, false)
		return tag.ptr
	}

	if currheapidx+uintptr(size) > heapendindex {
		panic("out of memory")
	}

	ansptr := currheapidx + 4
	write_tag(uint32(size), currheapidx, false)
	currheapidx += uintptr(size)
	write_tag(uint32(size), currheapidx-4, false)

	return ansptr
}

func Free(ptr uintptr) {
	szptr := ptr - 4
	sz := *(*uint32)(unsafe.Pointer(szptr))
	sz = sz & ^uint32(1)

	// left
	footerleftptr := szptr - 4
	if footerleftptr >= heapstrtindex && (*(*uint32)(unsafe.Pointer(footerleftptr))&1) == 1 {
		leftsz := *(*uint32)(unsafe.Pointer(footerleftptr)) & ^uint32(1)
		leftptr := szptr - uintptr(leftsz) + 4
		delete(freememory, ttags{ptr: leftptr, size: leftsz})
		sz += leftsz
		szptr -= uintptr(leftsz)
	}

	// right
	hi := szptr + uintptr(sz)
	headerrightptr := hi
	if headerrightptr < currheapidx && (*(*uint32)(unsafe.Pointer(headerrightptr))&1) == 1 {
		rightsz := *(*uint32)(unsafe.Pointer(headerrightptr)) & ^uint32(1)
		rightptr := headerrightptr + 4
		delete(freememory, ttags{ptr: rightptr, size: rightsz})
		sz += rightsz
	}

	write_tag(sz, szptr, true)
	write_tag(sz, szptr+uintptr(sz)-4, true)
	freememory[ttags{ptr: szptr + 4, size: sz}] = true
}
