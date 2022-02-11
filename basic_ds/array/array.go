package array

import (
	"fmt"
	"sync"
)

type Array struct {
	array []interface{}
	len   int
	cap   int
	lock  sync.Mutex
}

func NewArray(len int, cap int) *Array {
	s := new(Array)

	if len > cap {
		panic("Array's len is larger than cap")
	}

	array := make([]interface{}, cap, cap)

	s.array = array
	s.cap = cap
	s.len = 0

	return s
}

func (a *Array) Append(element interface{}) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if a.len == a.cap {
		newCap := 2 * a.len

		if a.cap == 0 {
			newCap = 1
		}

		newArray := make([]interface{}, newCap, newCap)

		for k, v := range a.array {
			newArray[k] = v
		}

		a.array = newArray
		a.cap = newCap
	}

	a.array[a.len] = element
	a.len = a.len + 1

}

func (a *Array) AppendMany(elements ...interface{}) {
	for _, v := range elements {
		a.Append(v)
	}
}

func (a *Array) Get(idx int) interface{} {
	a.lock.Lock()
	defer a.lock.Unlock()

	if a.len == 0 || idx >= a.len {
		panic("invalid index")
	}
	return a.array[idx]
}

func (a *Array) Cap() int {
	return a.cap
}

func (a *Array) Len() int {
	return a.len
}

func (a *Array) Print() {
	result := "["

	for i := 0; i < a.Len(); i++ {
		if i == 0 {
			result = fmt.Sprintf("%s%v", result, a.Get(i))
			continue
		}
		result = fmt.Sprintf("%s %v", result, a.Get(i))
	}
	result = result + "]"

	result = result + " " + fmt.Sprintf("cap: %d, len: %d", a.cap, a.len)

	fmt.Println(result)
}
