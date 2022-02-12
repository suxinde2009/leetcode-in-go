package tree

import (
	"fmt"
	"leetcode/basic_ds/queue"
)

type TreeNode struct {
	Data  interface{}
	Left  *TreeNode
	Right *TreeNode
}

func (treeNode *TreeNode) IsNil() bool {
	return treeNode == nil
}

func PreOder(tree *TreeNode) {

	if tree == nil {
		return
	}

	fmt.Printf("%v ", tree.Data)

	PreOder(tree.Left)

	PreOder(tree.Right)

}

func MidOder(tree *TreeNode) {

	if tree == nil {
		return
	}

	MidOder(tree.Left)

	fmt.Printf("%v ", tree.Data)

	MidOder(tree.Right)

}

func PostOder(tree *TreeNode) {

	if tree == nil {
		return
	}

	PostOder(tree.Left)

	PostOder(tree.Right)

	fmt.Printf("%v ", tree.Data)
}

func LayerOrder(treeNode *TreeNode) {

	if treeNode == nil {
		return
	}

	queue := new(queue.LinkQueue)
	queue.Add(treeNode)

	for queue.Size() > 0 {
		element, _ := queue.Remove().(*TreeNode)

		fmt.Printf("%v ", element.Data)

		if element.Left.IsNil() != true {
			queue.Add(element.Left)
		}

		if element.Right.IsNil() != true {
			queue.Add(element.Right)
		}
	}

}
