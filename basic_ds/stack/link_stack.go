package stack

import "sync"

type linkStackNode struct {
	Next  *linkStackNode
	Value interface{}
}

type LinkStack struct {
	root *linkStackNode
	size int
	lock sync.Mutex
}

func (stack *LinkStack) Push(v interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.root == nil {
		stack.root = new(linkStackNode)
		stack.root.Value = v
	} else {
		preNode := stack.root

		newNode := new(linkStackNode)
		newNode.Value = v

		newNode.Next = preNode

		stack.root = newNode
	}

	stack.size = stack.size + 1
}

func (stack *LinkStack) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.IsEmpty() {
		panic("Stack is empty, can't pop")
	}

	topNode := stack.root
	v := topNode.Value

	stack.root = topNode.Next

	stack.size = stack.size - 1

	return v
}

func (stack *LinkStack) Peek() interface{} {
	if stack.IsEmpty() {
		panic("Stack is empty, can't peek")
	}

	v := stack.root.Value
	return v
}

func (stack *LinkStack) Size() int {
	return stack.size
}

func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}
