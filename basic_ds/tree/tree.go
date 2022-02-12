package tree

import (
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

func PreOder(tree *TreeNode, visit func(value interface{})) {

	if tree == nil {
		return
	}

	visit(tree.Data)

	PreOder(tree.Left, visit)

	PreOder(tree.Right, visit)

}

func MidOder(tree *TreeNode, visit func(value interface{})) {

	if tree == nil {
		return
	}

	MidOder(tree.Left, visit)

	visit(tree.Data)

	MidOder(tree.Right, visit)

}

func PostOder(tree *TreeNode, visit func(value interface{})) {

	if tree == nil {
		return
	}

	PostOder(tree.Left, visit)

	PostOder(tree.Right, visit)

	visit(tree.Data)

}

func LayerOrder(treeNode *TreeNode, visit func(value interface{})) {

	if treeNode == nil {
		return
	}

	queue := new(queue.LinkQueue)
	queue.Add(treeNode)

	for queue.Size() > 0 {
		element, _ := queue.Remove().(*TreeNode)

		visit(element.Data)

		if element.Left.IsNil() != true {
			queue.Add(element.Left)
		}

		if element.Right.IsNil() != true {
			queue.Add(element.Right)
		}
	}

}
