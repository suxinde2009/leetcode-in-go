package tree

import (
	"fmt"
	"testing"
)

func If(isTrue bool, a string, b string) string {
	if isTrue {
		return a
	} else {
		return b
	}
}

func TestAVLTree(t *testing.T) {

	values := []int64{2, 3, 7, 10, 10, 10, 10, 23, 9, 102, 109, 111, 112, 113}

	tree := NewAVLTree()

	for _, v := range values {
		tree.Add(v)
	}

	fmt.Println("Find min value: ", tree.FindMinValue())
	fmt.Println("Find max value: ", tree.FindMaxValue())

	node := tree.Find(99)
	if node != nil {
		fmt.Println("Find it 99")
	} else {
		fmt.Println("Can not find 99")
	}

	node = tree.Find(9)
	if node != nil {
		fmt.Println("Find it 9")
	} else {
		fmt.Println("Can not find 9")
	}

	tree.Delete(9)
	tree.Delete(10)
	tree.Delete(2)
	tree.Delete(3)
	tree.Add(4)
	tree.Add(3)
	tree.Add(10)
	tree.Delete(111)

	node = tree.Find(9)
	if node != nil {
		fmt.Println("Find it 9")
	} else {
		fmt.Println("Can not find 9")
	}

	fmt.Println("\n-----------")
	tree.MidOder(func(value interface{}) {
		node, ok := value.(*AVLTreeNode)
		if !ok {
			return
		}
		fmt.Println(node)
	})

	fmt.Println("\n-----------")

	fmt.Println("The tree " + If(tree.IsAVLTree(), "is ", "is not ") + "an AVL tree")

}
