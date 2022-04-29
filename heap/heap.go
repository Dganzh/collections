package heap

type heap[T comparable] []T

func (h *heap[T]) push(elem T) {
	*h = append(*h, elem)
	h.heapify()
}

func (h *heap[T]) heapify() {

}
