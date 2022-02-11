package list

import (
	"testing"
)

func Test(t *testing.T) {
	node := NewLinkNode()
	node.Data = 2

	node1 := NewLinkNode()
	node1.Data = 3

	node.NextNode = node1

	node2 := NewLinkNode()
	node2.Data = 4

	node1.NextNode = node2

	PrintLinkList(node)
}
