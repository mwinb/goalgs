package main

import (
	"fmt"
	bst "goalgs/binarysearchtree"
	itd "goalgs/binarysearchtree/inttreedata"
)

func main() {
	fmt.Println("Go Algs")

	fmt.Println("\nBinary Search Tree")
	bst := bst.BinarySearchTree{}
	bst.Insert(itd.IntTreeData{Value: 10})
	intNode := bst.Find(10)

	fmt.Printf("\n %+v\n", intNode)
}
