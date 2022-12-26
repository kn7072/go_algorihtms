package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")

	root := &Node{Key: &Value{Value: 2}, Height: 1}
	insert(root, &Value{Value: 1})
	insert(root, &Value{Value: 3})
	insert(root, &Value{Value: 4})
	
	insert(root, &Value{Value: 5})
	insert(root, &Value{Value: 6})
	insert(root, &Value{Value: 7})
	insert(root, &Value{Value: 15})
	insert(root, &Value{Value: 16})

	insert(root, &Value{Value: 0})

	//deleteNode(root, nil, &Value{Value: 7})
	deleteNode(root, nil, &Value{Value: 6})
	
	fmt.Println()

	// res1 := root.search(NewData(13, 35))
	// fmt.Println(res1)

	// res2 := root.search(NewData(-1, 0))
	// fmt.Println(res2)

	// res3 := root.search(NewData(50, 55))
	// fmt.Println(res3)

	// fmt.Println()

}