package queue

import "sync"

type ArrayQueue struct {
	array []interface{}
	size  int
	lock  sync.Mutex
}

func (queue *ArrayQueue) Add(v interface{}) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	queue.array = append(queue.array, v)
	queue.size = queue.size + 1
}

func (queue *ArrayQueue) Remove() interface{} {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.IsEmpty() {
		panic("Queue is Empty, can't remove")
	}

	v := queue.array[0]

	newArray := make([]interface{}, queue.size-1, queue.size-1)
	for i := 1; i < queue.size; i++ {
		newArray[i-1] = queue.array[i]
	}

	queue.size = queue.size - 1

	return v
}

func (queue *ArrayQueue) Size() int {
	return queue.size
}

func (queue *ArrayQueue) IsEmpty() bool {
	return queue.size == 0
}
