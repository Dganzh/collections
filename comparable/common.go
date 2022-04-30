package comparable

import "golang.org/x/exp/constraints"

type Comparator[T comparable] func(T, T) int

func CommonComparator[T constraints.Ordered](a T, b T) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}
