package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := New[string]("")
	s.Push("1")
	fmt.Println(s.Top())
	fmt.Println(s.Pop())
	s.Push("2")
	s.Push("3")
	fmt.Println(s.Pop())

	s2 := New[int](0)
	s2.Push(1)
	fmt.Println(s.Top())
	fmt.Println(s.Pop())
	s2.Push(2)
	s2.Push(3)
	if s2.Pop() != 3 {
		t.Error("Pop error")
	}
}
