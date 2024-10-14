package algorithm

type numbers interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

type node[K comparable, V numbers] struct {
	key   K
	value V
}

type Heap[K comparable, V numbers] struct {
	nodes      []node[K, V]
	indexElem  map[K]int
	isContains map[K]bool
	size       int
}

func NewHeap[K comparable, V numbers]() (nodes Heap[K, V]) {
	heap := Heap[K, V]{
		nodes:      make([]node[K, V], 0),
		indexElem:  make(map[K]int),
		isContains: make(map[K]bool),
		size:       0,
	}

	return heap
}

func (heap *Heap[K, V]) Min() (key K, value V) {
	return heap.nodes[0].key, heap.nodes[0].value
}

func (heap *Heap[K, V]) ExtractMin() {
	heap.swap(0, len(heap.nodes)-1)
	heap.isContains[heap.nodes[len(heap.nodes)-1].key] = false
	heap.nodes = heap.nodes[:len(heap.nodes)-1]
	heap.size--
	heap.siftDown(0)
}

func (heap *Heap[K, V]) DecreaseKey(key K, newValue V) {
	index := heap.indexElem[key]
	heap.nodes[index].value = newValue
	heap.siftUp(index)
}

func (heap *Heap[K, V]) Insert(key K, value V) {
	heap.nodes = append(heap.nodes, node[K, V]{key, value})
	heap.isContains[key] = true
	heap.indexElem[key] = len(heap.nodes) - 1
	heap.siftUp(len(heap.nodes) - 1)

	heap.size++
}

func (heap *Heap[K, V]) Contains(key K) bool {
	value, ok := heap.isContains[key]
	return ok && value
}

func (heap *Heap[K, V]) Empty() bool {
	return heap.size == 0
}

func (heap *Heap[K, V]) siftUp(index int) {
	for index != 0 {
		parentIndex := heap.getParent(index)
		if heap.nodes[index].value >= heap.nodes[parentIndex].value {
			break
		}

		heap.swap(index, parentIndex)
		index = parentIndex
	}
}

func (heap *Heap[K, V]) siftDown(index int) {
	leftChild := heap.getLeftChild(index)
	rightChild := heap.getRightChild(index)

	if leftChild >= len(heap.nodes) {
		return
	}

	minIndex := index
	if leftChild < len(heap.nodes) &&
		heap.nodes[leftChild].value < heap.nodes[minIndex].value {
		minIndex = leftChild
	}

	if rightChild < len(heap.nodes) &&
		heap.nodes[rightChild].value < heap.nodes[minIndex].value {
		minIndex = rightChild
	}

	if minIndex != index {
		heap.swap(index, minIndex)
		heap.siftDown(minIndex)
	}
}

func (heap *Heap[K, V]) swap(i, j int) {
	heap.indexElem[heap.nodes[i].key] = j
	heap.indexElem[heap.nodes[j].key] = i
	heap.nodes[i], heap.nodes[j] = heap.nodes[j], heap.nodes[i]
}

func (heap *Heap[K, V]) getParent(index int) int {
	return (index - 1) / 2
}

func (heap *Heap[K, V]) getLeftChild(index int) int {
	return 2*index + 1
}

func (heap *Heap[K, V]) getRightChild(index int) int {
	return 2*index + 2
}
