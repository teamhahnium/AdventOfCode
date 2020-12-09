package main

var void = struct{}{}

// Set object..
type Set interface {
	Add(int)
	Contains(int) bool
}

type set struct {
	m map[int]struct{}
}

// NewSet creates a new Instace of Set
func NewSet() Set {
	return &set{
		m: make(map[int]struct{}),
	}
}

func (s *set) Add(v int) {
	s.m[v] = void
}

func (s *set) Contains(v int) bool {
	_, ok := s.m[v]
	return ok
}
