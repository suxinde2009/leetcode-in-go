package queue

import (
	"sync"
)

type linkQueueNode struct {
	Next  *linkQueueNode
	Value interface{}
}

type LinkQueue struct {
	root *linkQueueNode
	size int
	lock sync.Mutex
}

func (queue *LinkQueue) Add(v interface{}) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.root == nil {
		queue.root = new(linkQueueNode)
		queue.root.Value = v
	} else {

		newNode := new(linkQueueNode)
		newNode.Value = v

		nowNode := queue.root
		for nowNode.Next != nil {
			nowNode = nowNode.Next
		}
		nowNode.Next = newNode
	}
	queue.size = queue.size + 1
}

func (queue *LinkQueue) Remove() interface{} {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.IsEmpty() {
		panic("Queue is Empty, can't remove")
	}

	topNode := queue.root
	v := topNode.Value

	queue.root = topNode.Next

	queue.size = queue.size - 1

	return v
}

func (queue *LinkQueue) Size() int {
	return queue.size
}

func (queue *LinkQueue) IsEmpty() bool {
	return queue.size == 0
}
