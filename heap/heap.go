package heap

import cmp "collections/comparable"

type heap[T comparable] struct {
	limit      int
	comparator cmp.Comparator[T]
	nodes      []T
}

type Comparator[T any] func(T, T) int

func New[T comparable](limit int, defaultNode T, comparator cmp.Comparator[T]) *heap[T] {
	if limit <= 0 {
		limit = -1
	}
	h := &heap[T]{
		limit:      limit,
		comparator: comparator,
		nodes:      make([]T, 0, 1),
	}
	h.nodes = append(h.nodes, defaultNode)
	return h
}

func (h *heap[T]) Push(elem T) bool {
	if h.limit != -1 && h.Size() >= h.limit {
		return false
	}
	h.nodes = append(h.nodes, elem)
	h.up(len(h.nodes) - 1)
	return true
}

func (h *heap[T]) Pop() (elem T) {
	if h.Size() < 1 {
		return h.nodes[0]
	}
	elem = h.nodes[1]
	h.nodes[1] = h.nodes[len(h.nodes)-1]
	h.nodes = h.nodes[:len(h.nodes)-1]
	h.down(1)
	return
}

func (h *heap[T]) down(idx int) {
	var winner int
	for l := h.Left(idx); l <= h.Size(); l = h.Left(idx) {
		winner = h.winner(l, idx)
		r := h.Right(idx)
		if r <= h.Size() {
			winner = h.winner(winner, r)
		}
		if winner == idx {
			break
		}
		h.nodes[winner], h.nodes[idx] = h.nodes[idx], h.nodes[winner]
		idx = winner
	}
}

// 向上调整，维持堆性质
func (h *heap[T]) up(idx int) {
	if idx == 1 {
		return
	}
	for {
		p := h.Parent(idx)
		if h.winner(idx, p) == idx {
			h.nodes[idx], h.nodes[p] = h.nodes[p], h.nodes[idx]
		}
		if p == 1 {
			break
		}
		idx = p
	}
}

func (h *heap[T]) winner(idx1, idx2 int) int {
	if h.comparator(h.nodes[idx1], h.nodes[idx2]) > 0 {
		return idx1
	} else {
		return idx2
	}
}

func (h *heap[T]) Left(idx int) int {
	return idx << 1
}

func (h *heap[T]) Right(idx int) int {
	return idx<<1 | 1
}

func (h *heap[T]) Parent(idx int) int {
	return idx >> 1
}

func (h *heap[T]) Size() int {
	return len(h.nodes) - 1
}

func (h *heap[T]) Top() T {
	if len(h.nodes) > 1 {
		return h.nodes[1]
	}
	return h.nodes[0]
}

func (h *heap[T]) IsEmpty() bool {
	return len(h.nodes) == 1
}

func (h *heap[T]) IsFull() bool {
	if h.limit == -1 {
		return false
	}
	return len(h.nodes)-1 == h.limit
}

// Heapify 把传进的列表转成堆，如果传进的列表长度大于limit，将被无情截掉
func (h *heap[T]) Heapify(values []T) {
	if h.limit == -1 || len(values) <= h.limit {
		h.nodes = append(h.nodes[:1], values...)
		for i := len(values) / 2; i >= 1; i-- {
			h.down(i)
		}
	} else {
		h.Heapify(values[:h.limit])
	}
}
