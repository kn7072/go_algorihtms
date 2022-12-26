package main

import (
	"fmt"
)

func main(){
	root := NewTree()
	searchV := &Value{Key: 15}
	searchNotExist := &Value{Key: 50}
	Insert(&root, searchV)
	Insert(&root, &Value{Key: 10})
	Insert(&root, &Value{Key: 17})
	Insert(&root, &Value{Key: 5})
	Insert(&root, &Value{Key: 12})
	Insert(&root, &Value{Key: 30})
	Insert(&root, &Value{Key: 10})
	Insert(&root, &Value{Key: 31})
	Insert(&root, &Value{Key: 13})
	Insert(&root, &Value{Key: 14})

	res := Get(root, searchV)
	fmt.Println(res)
	
	res2 := Get(root, searchNotExist)
	fmt.Println(res2)
}