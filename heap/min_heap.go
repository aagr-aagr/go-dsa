package minheap

type kv struct {
	Key   int
	Value int
}

type MinHeap []kv

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Value < h[j].Value
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(kv))
}

func (h *MinHeap) Pop() any {
	i := len(*h) - 1
	v := (*h)[i]
	*h = (*h)[:i]
	return v
}

// Usage
//	h := make(MinHeap, 0, len(nums))
//	heap.Init(&h)
//  heap.Push(&h, kv{Key: 1, Value:5}
//}
