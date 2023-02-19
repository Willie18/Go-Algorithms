package hyperloglog

import (
	"fmt"
	"testing"
)

type hasher32 uint32

type hash32Test struct {
	hash uint32
}

func (h *hash32Test) Sum32() uint32 {
	return h.hash
}

func newHash32Test(x []byte) Hash32 {

	return &hash32Test{hash: uint32(len(x))}
}

func TestHyperLogLog(t *testing.T) {
	h := NewHyperLogLog(16)

	// Test adding items
	for i := 0; i < 100000; i++ {
		h.Add(newHash32Test([]byte(fmt.Sprintf("item%d", i))))
	}

	// Test estimating the count
	count := h.Count()
	if count < 90000 || count > 110000 {
		t.Errorf("Expected count to be between 90000 and 110000, got %f", count)
	}

	// Test adding items to a full HyperLogLog
	for i := 100000; i < 200000; i++ {
		h.Add(newHash32Test([]byte(fmt.Sprintf("item%d", i))))
	}
	count = h.Count()
	if count < 185000 || count > 215000 {
		t.Errorf("Expected count to be between 185000 and 215000, got %f", count)
	}
}
