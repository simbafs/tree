package aniavl

import (
	"fmt"

	"tree/tree"
)

type Node struct {
	value int
	left  *Node
	right *Node

	height int
	bf     int // balance factor
}

var _ tree.Node[Node] = &Node{}

func (n Node) View() string {
	return fmt.Sprintf("%d (h: %d, %d)", n.value, n.height, n.bf)
}

func (n Node) Children() []*Node {
	return []*Node{n.left, n.right}
}

func (n Node) IsNil() bool {
	return n.value == -1
}

var Nil = &Node{
	value:  -1,
	height: 0,
}

func NewNode(value int) *Node {
	return &Node{
		value:  value,
		height: 1,
		left:   Nil,
		right:  Nil,
	}
}
