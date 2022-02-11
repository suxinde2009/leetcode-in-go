package list

import "fmt"

type LinkNode struct {
	Data     interface{}
	NextNode *LinkNode
}

func NewLinkNode() *LinkNode {
	return &LinkNode{}
}

func PrintLinkList(node *LinkNode) {

	nowNode := node

	for {
		if nowNode != nil {
			fmt.Printf("%v\n", nowNode.Data)

			nowNode = nowNode.NextNode
			continue
		}
		break
	}

}
