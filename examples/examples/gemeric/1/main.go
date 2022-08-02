package main

import (
	//"constraints"
	//"fmt"
	"log"

	//"github.com/kn7072/go_algorihtms/constraints"
)

type Set[T comparable] map[T]struct{}

func Make[T comparable]() Set[T] {
    return make(Set[T])
}

func (s Set[T]) Add(v T) {
    s[v] = struct{}{}
}

func (s Set[T]) Delete(v T) {
    delete(s, v)
}

func (s Set[T]) Contains(v T) bool {
    _, ok := s[v]
    return ok
}

func (s Set[T]) Len() int {
    return len(s)
}

func (s Set[T]) Iterate(f func(T)) {
    for v := range s {
        f(v)
    }
}


func main() {
	s := Make[int]()
	s.Add(1)
	if s.Contains(2) {log.Fatalf("unexpected 2")}
}