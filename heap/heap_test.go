package heap

import (
	cmp "collections/comparable"
	xheap "container/heap"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"testing"
)

type HeapTestSuite struct {
	suite.Suite
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(HeapTestSuite))
}

type intSlice []int

func (s *intSlice) Len() int {
	return len(*s)
}

func (s *intSlice) Less(i, j int) bool {
	return (*s)[j] < (*s)[i]
}

func (s *intSlice) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *intSlice) Pop() (x any) {
	x = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}

func (s *intSlice) Push(x any) {
	*s = append(*s, x.(int))
}

func (s *HeapTestSuite) TestHeap() {
	h := &intSlice{1, 5, 4, 9}
	xheap.Init(h)
	myHeap := New[int](-1, 0, cmp.CommonComparator[int])
	myHeap.Heapify([]int{1, 5, 4, 9})
	s.Equalf([]int(*h), myHeap.nodes[1:], "heap values error")
	var rv int
	for i := 0; i < 10000; i++ {
		if h.Len() <= 0 {
			rv = 0
		} else {
			rv = rand.Intn(100)
		}
		if rv < 33 {
			v := rand.Intn(100000)
			xheap.Push(h, v)
			myHeap.Push(v)
		} else if rv < 66 {
			s.Equalf(xheap.Pop(h), myHeap.Pop(), "Pop error")
		} else if rv < 95 {
			s.Equalf((*h)[0], myHeap.Top(), "Top error")
		}
	}
}

type HeapCustomNode struct {
	num  float64
	name string
}

func MaxHeapCustomNodeComparator(n1 HeapCustomNode, n2 HeapCustomNode) int {
	if n1.num > n2.num {
		return 1
	} else if n1.num < n2.num {
		return -1
	} else {
		return 0
	}
}

func MinHeapCustomNodeComparator(n1 HeapCustomNode, n2 HeapCustomNode) int {
	if n2.num > n1.num {
		return 1
	} else if n2.num < n1.num {
		return -1
	} else {
		return 0
	}
}

func (s *HeapTestSuite) TestHeapCustomNode() {
	myHeap1 := New[HeapCustomNode](3, HeapCustomNode{}, MaxHeapCustomNodeComparator)
	myHeap1.Heapify([]HeapCustomNode{
		HeapCustomNode{10.0, "A"},
		HeapCustomNode{20.0, "B"},
		HeapCustomNode{30.0, "C"},
		HeapCustomNode{40.0, "D"},
	})
	s.Equal(3, myHeap1.Size(), "Size error")
	// 长度限制，第4个节点D被截掉，C是堆顶
	maxNode := myHeap1.Top()
	s.Equal(30.0, maxNode.num)
	s.Equal("C", maxNode.name)

	myHeap2 := New[HeapCustomNode](3, HeapCustomNode{}, MinHeapCustomNodeComparator)
	myHeap2.Heapify([]HeapCustomNode{
		HeapCustomNode{10.0, "A"},
		HeapCustomNode{20.0, "B"},
		HeapCustomNode{30.0, "C"},
		HeapCustomNode{40.0, "D"},
	})
	s.Equal(3, myHeap2.Size(), "Size error")
	// 长度限制，第4个节点D被截掉，A是堆顶
	minNode := myHeap2.Top()
	s.Equal(10.0, minNode.num)
	s.Equal("A", minNode.name)

	e := HeapCustomNode{9.0, "E"}
	// e应该直接被丢弃
	myHeap2.Push(e)
	s.Equal(3, myHeap2.Size(), "Size error")
	minNode = myHeap2.Top()
	s.Equal("A", minNode.name)

	// 先弹出
	myHeap2.Pop()
	s.Equal(2, myHeap2.Size(), "Size error")
	// 再推e进堆，应该成功
	myHeap2.Push(e)
	s.Equal(3, myHeap2.Size(), "Size error")
	minNode = myHeap2.Top()
	s.Equal("E", minNode.name)
}
