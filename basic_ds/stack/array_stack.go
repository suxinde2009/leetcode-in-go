package stack

import "sync"

type ArrayStack struct {
	array []interface{}

	size int
	lock sync.Mutex
}

func (stack *ArrayStack) Push(v interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	stack.array = append(stack.array, v)
	stack.size = stack.size + 1
}

func (stack *ArrayStack) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.IsEmpty() {
		panic("Stack is empty, can't pop")
	}

	v := stack.array[stack.size-1]

	newArray := make([]interface{}, stack.size-1, stack.size-1)
	for i := 0; i < stack.size-1; i++ {
		newArray[i] = stack.array[i]
	}
	stack.array = newArray

	stack.size = stack.size - 1

	return v
}

func (stack *ArrayStack) Peek() interface{} {
	if stack.IsEmpty() {
		panic("Stack is empty, can't peek")
	}

	v := stack.array[stack.size-1]

	return v
}

func (stack *ArrayStack) Size() int {
	return stack.size
}

func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}
