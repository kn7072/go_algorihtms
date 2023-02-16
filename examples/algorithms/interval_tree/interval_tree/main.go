package main

//https://iq.opengenus.org/interval-tree/

import (
	"fmt"
)

func main() {
	root := NewTree()
	searchV := NewValue(15, 16)
	searchNotExist := NewValue(50, 60)
	Insert(&root, searchV)
	Insert(&root, NewValue(10, 11))
	Insert(&root, NewValue(17, 18))
	Insert(&root, NewValue(5, 6))
	Insert(&root, NewValue(12, 13))
	Insert(&root, NewValue(30, 31))
	Insert(&root, NewValue(10, 12))
	Insert(&root, NewValue(31, 32))
	Insert(&root, NewValue(13, 14))
	Insert(&root, NewValue(14, 16))

	res := Get(root, searchV)
	fmt.Println(res)

	res2 := Get(root, searchNotExist)
	fmt.Println(res2)

	r := overlap(root, NewValue(10, 12))
	fmt.Println(r)

	r = overlap(root, NewValue(14, 15))
	fmt.Println(r)

	r = overlap(root, NewValue(33, 34))
	fmt.Println(r)

	r = overlap(root, NewValue(31, 31))
	fmt.Println(r)

	r = overlap(root, NewValue(31, 32))
	fmt.Println(r)

	r = overlap(root, NewValue(10, 10))
	fmt.Println(r)

	r = overlap(root, NewValue(5, 6))
	fmt.Println(r)

	// Delete(&root, NewValue(5, 6))
	// Delete(&root, NewValue(12, 13))
	Delete(&root, NewValue(15, 16))

	fmt.Println()
}
