package tree

func nop() {}

func max(l int64, r int64) int64 {
	if l >= r {
		return l
	}
	return r
}

type AVLTreeNode struct {
	Value  int64
	Times  int64
	Height int64
	Left   *AVLTreeNode
	Right  *AVLTreeNode
}

func NewAVLTree() *AVLTree {
	return new(AVLTree)
}

type AVLTree struct {
	Root *AVLTreeNode
}

func (node *AVLTreeNode) UpdateHeight() {
	if node == nil {
		return
	}

	var leftHeight, rightHeight int64 = 0, 0

	if node.Left != nil {
		leftHeight = node.Left.Height
	}
	if node.Right != nil {
		rightHeight = node.Right.Height
	}

	maxHeight := max(leftHeight, rightHeight)

	node.Height = maxHeight + 1
}

func (node *AVLTreeNode) BalanceFactor() int64 {
	var leftHeight, rightHeight int64 = 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}
	if node.Right != nil {
		rightHeight = node.Right.Height
	}
	return leftHeight - rightHeight
}

func RightRotation(Root *AVLTreeNode) *AVLTreeNode {
	Piovt := Root.Left
	B := Piovt.Right

	Piovt.Right = Root
	Root.Left = B

	Root.UpdateHeight()
	Piovt.UpdateHeight()
	return Piovt
}

func LeftRoration(Root *AVLTreeNode) *AVLTreeNode {
	Pivot := Root.Right
	B := Pivot.Left

	Pivot.Left = Root
	Root.Right = B

	Root.UpdateHeight()
	Pivot.UpdateHeight()
	return Pivot
}

func LeftRightRoration(node *AVLTreeNode) *AVLTreeNode {
	node.Left = LeftRoration(node.Left)
	return RightRotation(node)
}

func RightLeftRoration(node *AVLTreeNode) *AVLTreeNode {
	node.Right = RightRotation(node.Right)
	return LeftRoration(node)
}

func (tree *AVLTree) Add(value int64) {
	tree.Root = tree.Root.Add(value)
}

func (node *AVLTreeNode) Add(value int64) *AVLTreeNode {
	if node == nil {
		return &AVLTreeNode{Value: value, Height: 1}
	}

	if node.Value == value {
		node.Times += 1
		return node
	}

	var newTreeNode *AVLTreeNode
	if value > node.Value {

		node.Right = node.Right.Add(value)
		factor := node.BalanceFactor()

		if factor == -2 {
			if value > node.Right.Value {
				newTreeNode = LeftRoration(node)
			} else {
				newTreeNode = RightLeftRoration(node)
			}
		} else {
			nop()
		}

	} else {
		node.Left = node.Left.Add(value)
		factor := node.BalanceFactor()

		if factor == 2 {
			if value < node.Left.Value {
				newTreeNode = RightRotation(node)
			} else {
				newTreeNode = LeftRightRoration(node)
			}
		} else {
			nop()
		}
	}

	if newTreeNode == nil {
		node.UpdateHeight()
		return node
	} else {
		newTreeNode.UpdateHeight()
		return newTreeNode
	}

}

func (tree *AVLTree) FindMinValue() *AVLTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindMinValue()
}

func (node *AVLTreeNode) FindMinValue() *AVLTreeNode {
	if node.Left == nil {
		return node
	}

	return node.Left.FindMinValue()
}

func (tree *AVLTree) FindMaxValue() *AVLTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindMaxValue()
}

func (node *AVLTreeNode) FindMaxValue() *AVLTreeNode {
	if node.Right == nil {
		return node
	}

	return node.Right.FindMaxValue()
}

func (tree *AVLTree) Find(value int64) *AVLTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.Find(value)
}

func (node *AVLTreeNode) Find(value int64) *AVLTreeNode {
	if node.Value == value {
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

func (tree *AVLTree) MidOder(visit func(value interface{})) {
	tree.Root.MidOrder(visit)
}

func (node *AVLTreeNode) MidOrder(visit func(value interface{})) {
	if node == nil {
		return
	}

	node.Left.MidOrder(visit)

	for i := 0; i <= int(node.Times); i++ {
		visit(node)
	}

	node.Right.MidOrder(visit)
}

func (tree *AVLTree) Delete(value int64) {
	if tree.Root == nil {
		return
	}
	tree.Root = tree.Root.Delete(value)
}

func (node *AVLTreeNode) Delete(value int64) *AVLTreeNode {
	if node == nil {
		return nil
	}
	if value < node.Value {
		node.Left = node.Left.Delete(value)
		node.Left.UpdateHeight()
	} else if value > node.Value {
		node.Right = node.Right.Delete(value)
		node.Right.UpdateHeight()
	} else {
		if node.Left == nil && node.Right == nil {
			return nil
		}
		if node.Left != nil && node.Right != nil {
			if node.Left.Height > node.Right.Height {
				maxNode := node.Left
				for maxNode.Right != nil {
					maxNode = maxNode.Right
				}
				node.Value = maxNode.Value
				node.Times = maxNode.Times

				node.Left = node.Left.Delete(maxNode.Value)
				node.Left.UpdateHeight()
			} else {

				minNode := node.Right
				for minNode.Left != nil {
					minNode = minNode.Left
				}

				node.Value = minNode.Value
				node.Times = minNode.Times

				node.Right = node.Right.Delete(minNode.Value)
				node.Right.UpdateHeight()

			}
		} else {

			if node.Left != nil {

				node.Value = node.Left.Value
				node.Times = node.Left.Times
				node.Height = 1
				node.Left = nil

			} else { //if node.Right != nil {

				node.Value = node.Right.Value
				node.Times = node.Right.Times
				node.Height = 1
				node.Right = nil

			}

		}
		return node
	}

	var newNode *AVLTreeNode
	if node.BalanceFactor() == 2 {

		if node.Left.BalanceFactor() >= 0 {
			newNode = RightRotation(node)
		} else {
			newNode = LeftRightRoration(node)
		}

	} else if node.BalanceFactor() == -2 {
		if node.Right.BalanceFactor() >= 0 {
			newNode = LeftRoration(node)
		} else {
			newNode = RightLeftRoration(node)
		}

	} else {
		nop()
	}

	if newNode == nil {
		node.UpdateHeight()
		return node
	} else {
		newNode.UpdateHeight()
		return newNode
	}
}

func (tree *AVLTree) IsAVLTree() bool {
	if tree == nil || tree.Root == nil {
		return true
	}
	return tree.Root.IsRight()
}

func (node *AVLTreeNode) IsRight() bool {
	if node == nil {
		return true
	}

	if node.Left == nil && node.Right == nil {
		if node.Height == 1 {
			return true
		} else {
			return false
		}
	} else if node.Left != nil && node.Right != nil {

		if node.Left.Value < node.Value && node.Right.Value > node.Value {
			nop()
		} else {
			return false
		}

		bal := node.Left.Height - node.Right.Height
		if bal < 0 {
			bal = -bal
		}

		if bal > 1 {
			return false
		}

		if node.Left.Height > node.Right.Height {
			if node.Height == node.Left.Height+1 {
				nop()
			} else {
				return false
			}
		} else {
			if node.Height == node.Right.Height+1 {
				nop()
			} else {
				return false
			}
		}

		if !node.Left.IsRight() {
			return false
		}

		if !node.Right.IsRight() {
			return false
		}

	} else {

		if node.Right != nil {
			if node.Right.Height == 1 && node.Right.Left == nil && node.Right.Right == nil {
				if node.Right.Value > node.Value {
					nop()
				} else {
					return false
				}
			} else {
				return false
			}
		} else {
			if node.Left.Height == 1 && node.Left.Left == nil && node.Left.Right == nil {
				if node.Left.Value < node.Value {
					nop()
				} else {
					return false
				}
			} else {
				return false
			}
		}

	}

	return true
}
