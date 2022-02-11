package ring

import (
	"fmt"
)

type Ring struct {
	next, prev *Ring

	Value interface{}
}

// O(1)
func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

// O(1)
func NewRing() *Ring {
	return &Ring{}
}

// O(n)
func CreateRing(ringNodeCnt int) *Ring {
	if ringNodeCnt <= 0 {
		return nil
	}
	r := NewRing()
	p := r

	for i := 1; i < ringNodeCnt; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}

	p.next = r
	r.prev = p
	return r
}

// O(1)
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

// O(1)
func (r *Ring) Prev() *Ring {
	if r.prev == nil {
		return r.init()
	}
	return r.prev
}

// O(n)
func (r *Ring) Move(n int) *Ring {

	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

// O(n)
func (r *Ring) Link(s *Ring) *Ring {
	if s == nil {
		return r
	}
	n := r.Next()

	p := s.Prev()
	r.next = s
	s.prev = r
	n.prev = p
	p.next = n

	return n
}

// O(n)
func (r *Ring) Unlink(nodeCnt int) *Ring {
	if nodeCnt < 0 {
		return nil
	}
	return r.Link(r.Move(nodeCnt + 1))
}

func (r *Ring) Len() int {
	n := 0

	if r != nil {
		n = 1

		for p := r.Next(); p != r; p = p.next {
			n++
		}
	}
	return n
}

func PrintRing(r *Ring) {
	node := r

	for {
		fmt.Printf("%v\n", node.Value)
		node = node.Next()

		if node == r {
			break
		}
	}
}
