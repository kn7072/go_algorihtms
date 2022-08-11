package main

import (
	"fmt"
)

//https://habr.com/ru/company/skillfactory/blog/657853/
// https://go.googlesource.com/proposal/+/HEAD/design/43651-type-parameters.md

type Tree[T interface{}] struct {
    left, right *Tree[T]
    value       T
}

func (t *Tree[T]) Lookup(x T) *Tree[T] {
	fmt.Println(t.value)
	fmt.Printf("type of t %T, %v\n", t, t)
	return nil //t.left
}

func NewTree[T interface{}] (value T) Tree[T] {
	return Tree[T]{value: value,
				   left: nil,
				   right: nil,
				  }
}

var stringTree = Tree[string]{value: "11"}
/*
Tree хранятся значения параметра типа T. В дженерик-типах могут быть и методы, такие как Lookup выше. 
Чтобы использовать дженерик-тип, его нужно инстанцировать. 
Tree[string] — пример инстанцирования Tree с типом-аргументом string.
*/

/*
type Ordered interface {
    Integer|Float|~string
}

Integer и Float — это интерфейсные типы, аналогично определённых в пакете constraints. 
Обратите внимание: нет методов, определяемых интерфейсом Ordered.

Что касается ограничений типа, конкретный тип (например, string) нас обычно не так интересует, 
как все строковые типы. Вот для чего нужен токен ~: выражение ~string означает набор всех типов с базовым типом string. 
Это сам тип string и все типы, объявленные с такими определениями, как type MyString string.



// Scale returns a copy of s with each element multiplied by c.
func Scale[S ~[]E, E constraints.Integer](s S, c E) S {
    r := make(S, len(s))
    for i, v := range s {
        r[i] = v * c
    }
    return r
}

*/

func main() {
	//testObject := Tree[string]{}
	stringTree.Lookup("testObject")
}