package bst

import (
	"fmt"

	"tree/tree"
)

type Node struct {
	Key   int
	Left  *Node
	Right *Node
	isNil bool
}

var _ tree.Node = &Node{}

var Nil = &Node{
	Key:   -1,
	isNil: true,
}

func NewNode(key int) *Node {
	return &Node{
		Key:   key,
		Left:  Nil,
		Right: Nil,
	}
}

// common methods for Node

func (node Node) View() string {
	return fmt.Sprintf("%d", node.Key)
}

func (node Node) Children() []tree.Node {
	return []tree.Node{node.Left, node.Right}
}

func (node Node) IsNil() bool {
	return node.isNil
}

// custom methods for BSTNode

func (node *Node) Insert(newNode *Node) {
	if newNode.Key < node.Key {
		if node.Left.IsNil() {
			node.Left = newNode
		} else {
			node.Left.Insert(newNode)
		}
	} else {
		if node.Right.IsNil() {
			node.Right = newNode
		} else {
			node.Right.Insert(newNode)
		}
	}
}
