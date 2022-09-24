package main

import (
	"fmt"
)

func main() {
	uf := NewUnionFind(10)

	// Union a,b connects components at index a and b
	uf.Union(1, 2)
	uf.Union(2, 3)
	uf.Union(5, 6)
	uf.Union(4, 6)

	fmt.Println(uf.Find(2)) // Prints 1
	fmt.Println(uf.Find(3)) // Prints 1
	fmt.Println(uf.Connected(1, 2)) // Prints true
	fmt.Println(uf.Connected(1, 6)) // Prints false
}