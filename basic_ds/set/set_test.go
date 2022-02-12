package set

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	s := NewSet(5)

	s.Add(1)
	s.Add(2)
	s.Add(1)

	fmt.Println("list of all items in set: ", s.List())

	s.Clear()
	if s.IsEmpty() {
		fmt.Println("The set is cleared, now is empty")
	}

	s.Add(1)
	s.Add(2)
	s.Add(3)

	if s.Has(2) {
		fmt.Println("2 exists in the set")
	}

	s.Remove(2)
	s.Remove(3)
	fmt.Println("list of all items in set: ", s.List())
}
