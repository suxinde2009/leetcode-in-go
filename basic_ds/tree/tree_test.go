package tree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {

	tree := &TreeNode{Data: "A"}
	tree.Left = &TreeNode{Data: "B"}
	tree.Right = &TreeNode{Data: "C"}

	tree.Left.Left = &TreeNode{Data: "D"}
	tree.Left.Right = &TreeNode{Data: "E"}

	tree.Right.Left = &TreeNode{Data: "F"}

	fmt.Println("PreOrder ")
	PreOder(tree, func(value interface{}) {
		fmt.Printf("%v ", value)
	})
	fmt.Println()

	fmt.Println("MidOrder ")
	MidOder(tree, func(value interface{}) {
		fmt.Printf("%v ", value)
	})
	fmt.Println()

	fmt.Println("PostOrder ")
	PostOder(tree, func(value interface{}) {
		fmt.Printf("%v ", value)
	})
	fmt.Println()

	fmt.Println("LayerOrder ")
	LayerOrder(tree, func(value interface{}) {
		fmt.Printf("%v ", value)
	})

	fmt.Println()
}
