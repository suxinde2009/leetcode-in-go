package queue

type IQueue interface {
	Add(v interface{})
	Remove() interface{}
	Size() int
	IsEmpty() bool
}
