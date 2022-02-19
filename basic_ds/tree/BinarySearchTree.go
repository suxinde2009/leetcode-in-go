package tree

type BinarySearchTreeNode struct {
	Value int64
	Times int64
	Left  *BinarySearchTreeNode
	Right *BinarySearchTreeNode
}

type BinarySearchTree struct {
	Root *BinarySearchTreeNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return new(BinarySearchTree)
}

func (tree *BinarySearchTree) Add(value int64) {
	if tree.Root == nil {
		tree.Root = &BinarySearchTreeNode{Value: value}
		return
	}
	tree.Root.Add(value)
}

func (node *BinarySearchTreeNode) Add(value int64) {
	if value < node.Value {

		if node.Left == nil {
			node.Left = &BinarySearchTreeNode{Value: value}
		} else {
			node.Left.Add(value)
		}

	} else if value > node.Value {
		if node.Right == nil {
			node.Right = &BinarySearchTreeNode{Value: value}
		} else {
			node.Right.Add(value)
		}
	} else {
		node.Times = node.Times + 1
	}
}

func (tree *BinarySearchTree) FindMinValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindMinValue()
}

func (node *BinarySearchTreeNode) FindMinValue() *BinarySearchTreeNode {
	if node.Left == nil {
		return node
	}
	return node.Left.FindMinValue()
}

func (tree *BinarySearchTree) FindMaxValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindMaxValue()
}

func (node *BinarySearchTreeNode) FindMaxValue() *BinarySearchTreeNode {
	if node.Right == nil {
		return node
	}
	return node.Right.FindMaxValue()
}

func (tree *BinarySearchTree) Find(value int64) *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}

	return tree.Root.Find(value)
}

func (node *BinarySearchTreeNode) Find(value int64) *BinarySearchTreeNode {
	if value == node.Value {
		return node
	} else if value < node.Value {
		if node.Left == nil {
			return nil
		}
		return node.Left.Find(value)
	} else {
		if node.Right == nil {
			return nil
		}
		return node.Right.Find(value)
	}
}

func (tree *BinarySearchTree) FindParent(value int64) *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	if tree.Root.Value == value {
		return nil
	}

	return tree.Root.FindParent(value)
}

func (node *BinarySearchTreeNode) FindParent(value int64) *BinarySearchTreeNode {
	if value < node.Value {

		leftTree := node.Left

		if leftTree == nil {
			return nil
		}

		if leftTree.Value == value {
			return node

		} else {
			return leftTree.FindParent(value)
		}

	} else {

		rightTree := node.Right
		if rightTree == nil {
			return nil
		}
		if rightTree.Value == value {
			return node
		} else {
			return rightTree.FindParent(value)
		}
	}

}

func (tree *BinarySearchTree) Delete(value int64) {
	if tree.Root == nil {
		return
	}

	node := tree.Root.Find(value)
	if node == nil {
		return
	}

	parent := tree.Root.FindParent(value)

	if parent == nil && node.Left == nil && node.Right == nil {
		tree.Root = nil
		return
	} else if node.Left == nil && node.Right == nil {
		if parent.Left != nil && value == parent.Left.Value {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		return
	} else if node.Left != nil && node.Right != nil {

		minNode := node.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		tree.Delete(minNode.Value)

		node.Value = minNode.Value
		node.Times = minNode.Times

	} else {

		if parent == nil {
			if node.Left != nil {
				tree.Root = node.Left
			} else {
				tree.Root = node.Right
			}
			return
		}
		if node.Left != nil {
			if parent.Left != nil && value == parent.Left.Value {
				parent.Left = node.Left
			} else {
				parent.Right = node.Left
			}
		} else {
			if parent.Left != nil && value == parent.Left.Value {
				parent.Left = node.Right
			} else {
				parent.Right = node.Right
			}
		}

	}

}

func (tree *BinarySearchTree) MidOder(visit func(value interface{})) {
	tree.Root.MidOrder(visit)
}

func (node *BinarySearchTreeNode) MidOrder(visit func(value interface{})) {
	if node == nil {
		return
	}

	node.Left.MidOrder(visit)

	for i := 0; i <= int(node.Times); i++ {
		visit(node)
	}

	node.Right.MidOrder(visit)
}
