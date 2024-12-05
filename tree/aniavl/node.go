package aniavl

import (
	"fmt"

	"tree/tree"

	"github.com/charmbracelet/lipgloss"
)

type Node struct {
	value int
	left  *Node
	right *Node

	height int
	bf     int // balance factor

	isCurr bool
}

var _ tree.Node[Node] = &Node{}

func (n Node) View() string {
	s := fmt.Sprintf("%d (h: %d, %d)", n.value, n.height, n.bf)
	if n.isCurr {
		return lipgloss.NewStyle().
			Background(lipgloss.Color("#cccccc")).
			Foreground(lipgloss.Color("#000000")).
			Render(s)
	} else {
		return s
	}
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

func (n *Node) UpdateHeight() {
	n.height = 1 + max(n.left.height, n.right.height)
	n.bf = n.right.height - n.left.height
}

func (n *Node) SetLeft(left *Node) {
	n.left = left
}

func (n *Node) SetRight(right *Node) {
	n.right = right
}

func (n *Node) Inorder() []int {
	if n.IsNil() {
		return []int{}
	}

	return append(append(n.left.Inorder(), n.value), n.right.Inorder()...)
}
