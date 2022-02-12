package stack

import (
	"fmt"
	"testing"
)

func TestArrayStack(t *testing.T) {
	fmt.Println("TestArrayStack =================")

	s := new(ArrayStack)
	s.Push("cat")
	s.Push("dog")
	s.Push("hen")
	fmt.Println("size: ", s.Size())
	fmt.Println("pop: ", s.Pop())
	fmt.Println("pop: ", s.Pop())
	fmt.Println("size: ", s.Size())
	s.Push("drag")
	fmt.Println("size: ", s.Size())
	fmt.Println("peek: ", s.Peek())
	fmt.Println("size: ", s.Size())

	fmt.Println("IsEmpty: ", s.IsEmpty())
	s.Pop()
	s.Pop()
	fmt.Println("IsEmpty: ", s.IsEmpty())
}

func TestLinkStack(t *testing.T) {
	fmt.Println("TestLinkStack =================")

	s := new(LinkStack)
	s.Push("cat")
	s.Push("dog")
	s.Push("hen")
	fmt.Println("size: ", s.Size())
	fmt.Println("pop: ", s.Pop())
	fmt.Println("pop: ", s.Pop())
	fmt.Println("size: ", s.Size())
	s.Push("drag")
	fmt.Println("size: ", s.Size())
	fmt.Println("peek: ", s.Peek())
	fmt.Println("size: ", s.Size())

	fmt.Println("IsEmpty: ", s.IsEmpty())
	s.Pop()
	s.Pop()
	fmt.Println("IsEmpty: ", s.IsEmpty())
}
