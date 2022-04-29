package stack

type _stack[T comparable] struct {
	values   []T
	_default T
}

func New[T comparable](defaultVal T) *_stack[T] {
	return &_stack[T]{
		values:   []T{},
		_default: defaultVal,
	}
}

func (s *_stack[T]) Push(elem T) {
	s.values = append(s.values, elem)
}

func (s *_stack[T]) Pop() (elem T) {
	if len(s.values) > 0 {
		elem = s.values[len(s.values)-1]
		s.values = s.values[:len(s.values)-1]
		return
	}
	return s._default
}

func (s *_stack[T]) Top() T {
	if len(s.values) > 0 {
		return s.values[len(s.values)-1]
	}
	return s._default
}
