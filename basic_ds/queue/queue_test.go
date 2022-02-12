package queue

import (
	"fmt"
	"testing"
)

func TestArrayQueue(t *testing.T) {
	fmt.Println("TestArrayQueue =================")

	q := new(ArrayQueue)
	q.Add("cat")
	q.Add("dog")
	q.Add("hen")
	fmt.Println("size: ", q.Size())
	fmt.Println("remove: ", q.Remove())
	fmt.Println("remove: ", q.Remove())
	fmt.Println("size: ", q.Size())
	q.Add("drag")
	fmt.Println("size: ", q.Size())
	fmt.Println("remove: ", q.Remove())
	fmt.Println("size: ", q.Size())

	fmt.Println("IsEmpty: ", q.IsEmpty())
	q.Remove()
	fmt.Println("IsEmpty: ", q.IsEmpty())
}

func TestLinkQueue(t *testing.T) {
	fmt.Println("TestLinkQueue =================")

	q := new(LinkQueue)
	q.Add("cat")
	q.Add("dog")
	q.Add("hen")
	fmt.Println("size: ", q.Size())
	fmt.Println("remove: ", q.Remove())
	fmt.Println("remove: ", q.Remove())
	fmt.Println("size: ", q.Size())
	q.Add("drag")
	fmt.Println("size: ", q.Size())
	fmt.Println("remove: ", q.Remove())
	fmt.Println("size: ", q.Size())

	fmt.Println("IsEmpty: ", q.IsEmpty())
	q.Remove()
	fmt.Println("IsEmpty: ", q.IsEmpty())
}
