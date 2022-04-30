package main

import (
	"fmt"
)

type Node struct {
	num int
}

type Number[T ~int | ~int8] T

func CompareTo(x Number, y Number) bool {
	return x > y
}

func main() {
	n := Node{10}
	i := 10
	fmt.Println(n.num == i)
}
