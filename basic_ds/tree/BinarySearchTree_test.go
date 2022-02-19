package tree

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {

	values := []int64{3, 6, 8, 30, 9, 2, 6, 8, 9, 3, 5, 60, 7, 8, 13, 6, 8}

	tree := NewBinarySearchTree()
	for _, v := range values {
		tree.Add(v)
	}

	fmt.Println("Find min value: ", tree.FindMinValue())
	fmt.Println("Find max value: ", tree.FindMaxValue())

	node := tree.Find(99)
	if node != nil {
		fmt.Println("Found element 99")
	} else {
		fmt.Println("Can not find element 99")
	}

	node = tree.Find(2)
	if node != nil {
		fmt.Println("Found element 2")
	} else {
		fmt.Println("Can not find element 2")
	}

	tree.Delete(2)
	node = tree.Find(2)
	if node != nil {
		fmt.Println("Found element 2")
	} else {
		fmt.Println("Can not find element 2")
	}

	fmt.Println("----")

	tree.MidOder(func(value interface{}) {

		fmt.Println(value)
	})

	fmt.Println("\n----")

}
