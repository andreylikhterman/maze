package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHeap_EmptyHeap(t *testing.T) {
	heap := NewHeap[int, int]()

	assert.True(t, heap.Empty())
}

func TestHeap_Insert_And_Contains(t *testing.T) {
	heap := NewHeap[int, int]()
	heap.Insert(1, 10)
	heap.Insert(2, 5)
	heap.Insert(3, 15)

	assert.False(t, heap.Empty())
	assert.True(t, heap.Contains(1))
	assert.True(t, heap.Contains(2))
	assert.True(t, heap.Contains(3))
	assert.False(t, heap.Contains(4))
}

func TestHeap_Min(t *testing.T) {
	heap := NewHeap[int, int]()
	heap.Insert(1, 10)
	heap.Insert(2, 5)
	heap.Insert(3, 15)

	key, value := heap.Min()
	assert.Equal(t, 2, key)
	assert.Equal(t, 5, value)
}

func TestHeap_ExtractMin(t *testing.T) {
	heap := NewHeap[int, int]()
	heap.Insert(1, 10)
	heap.Insert(2, 5)
	heap.Insert(3, 15)

	heap.ExtractMin()

	assert.Equal(t, 2, heap.size)
	key, _ := heap.Min()
	assert.Equal(t, 1, key)
}

func TestHeap_DecreaseKey(t *testing.T) {
	heap := NewHeap[int, int]()
	heap.Insert(1, 10)
	heap.Insert(2, 5)
	heap.Insert(3, 15)

	heap.DecreaseKey(3, 2)

	key, _ := heap.Min()
	assert.Equal(t, 3, key)
}

func TestHeap_EmptyHeapAfterExtraction(t *testing.T) {
	heap := NewHeap[int, int]()
	heap.Insert(1, 10)
	heap.Insert(2, 5)

	heap.ExtractMin()
	heap.ExtractMin()

	assert.True(t, heap.Empty())
}
