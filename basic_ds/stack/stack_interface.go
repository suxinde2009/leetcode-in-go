package stack

type IStack interface {
	Push(v interface{})
	Pop() interface{}
	Peek() interface{}
	Size() int
	IsEmpty() bool
}
