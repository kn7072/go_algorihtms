package main

import (
	"fmt"
	"sync"
)

// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#embedded-type-parameter
// A Lockable is a value that may be safely simultaneously accessed
// from multiple goroutines via the Get and Set methods.
type Lockable[T any] struct {
	v T
	mu sync.Mutex
}

// Get returns the value stored in a Lockable.
func (l *Lockable[T]) Get() T {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.v
}

// Set sets the value in a Lockable.
func (l *Lockable[T]) Set(v T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.v = v
}

// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#embedded-type-parameter-methods
// NamedInt is an int with a name. The name can be any type with
// a String method.

type MyStringer struct {
	name string
}

func (s MyStringer) String() string {
	return s.name
}

type NamedInt[Name fmt.Stringer] struct {
	name Name
	val int
}

// Name returns the name of a NamedInt.
func (ni NamedInt[Name]) Name() string {
	// The String method is promoted from the embedded Name.
	return ni.name.String()
}

func main() {
	l := Lockable[int]{v: 2}
	g := l.Get()
	fmt.Println(g)

	str := NamedInt[MyStringer]{name: MyStringer{"123"}}
	s := str.Name()
	fmt.Println(s)
}