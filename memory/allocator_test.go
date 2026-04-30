package memory

import (
	"testing"
)

func TestAllocator(t *testing.T) {
	initheap(int(N))
	
	// Reset global state
	freememory = make(map[ttags]bool)
	currheapidx = heapstrtindex

	// 1. Basic Malloc/Free
	p1 := Malloc(100)
	t.Logf("p1: %x", p1)
	if p1 == 0 {
		t.Fatal("Malloc(100) failed")
	}
	Free(p1)
	
	// 2. Reuse
	p2 := Malloc(100)
	t.Logf("p2: %x", p2)
	if p2 != p1 {
		t.Errorf("Expected reuse of p1 (%x), got %x", p1, p2)
	}
	Free(p2)

	// 3. Coalescing
	p3 := Malloc(100)
	t.Logf("p3: %x", p3)
	p4 := Malloc(100)
	t.Logf("p4: %x", p4)
	p5 := Malloc(100)
	t.Logf("p5: %x", p5)
	_ = p5 // avoid unused

	Free(p3)
	Free(p4) // Should coalesce with p3

	p6 := Malloc(200)
	t.Logf("p6: %x", p6)
	if p6 != p3 {
		t.Errorf("Expected coalesced reuse at %x, got %x", p3, p6)
	}

	// 4. Triple Coalescing
	p7 := Malloc(100)
	t.Logf("p7: %x", p7)
	p8 := Malloc(100)
	t.Logf("p8: %x", p8)
	p9 := Malloc(100)
	t.Logf("p9: %x", p9)
	p10 := Malloc(100)
	t.Logf("p10: %x", p10)
	_ = p10

	Free(p7)
	Free(p9)
	Free(p8) // Should join p7, p8, and p9

	p11 := Malloc(300)
	t.Logf("p11: %x", p11)
	if p11 != p7 {
		t.Errorf("Expected triple coalesced reuse at %x, got %x", p7, p11)
	}

	// 5. Splitting
	p12 := Malloc(500)
	t.Logf("p12: %x", p12)
	Free(p12) // Creates 512 byte block

	p13 := Malloc(200) // Needs 208
	t.Logf("p13: %x", p13)
	if p13 != p12 {
		t.Errorf("Expected splitting to use start of p12 (%x), got %x", p12, p13)
	}

	p14 := Malloc(200) // Needs 208. Fits in remainder (512 - 208 = 304)
	t.Logf("p14: %x", p14)
	if p14 != p12+208 {
		t.Errorf("Expected p14 to take remainder of p12 at %x, got %x", p12+208, p14)
	}
}
