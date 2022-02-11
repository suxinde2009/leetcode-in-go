package stack

import "sync"

type ArrayStack struct {
	array []interface{}

	size int
	lock sync.Mutex
}
