package avl

import (
	"fmt"
	"log"

	"tree/tree"
)

type Node struct {
	Key   int
	Left  *Node
	Right *Node

	Height int
	BF     int // balance factor
}

var _ tree.Node[Node] = &Node{}

// common methods for Node

func (node Node) View() string {
	return fmt.Sprintf("%d (h: %d, %d)", node.Key, node.Height, node.BF)
}

func (node Node) Children() []*Node {
	return []*Node{node.Left, node.Right}
}

func (node Node) IsNil() bool {
	return node.Key == -1
}

// custom methods for Node

var Nil = &Node{
	Key:    -1,
	Height: 0,
}

func NewNode(key int) *Node {
	return &Node{
		Key:    key,
		Height: 1,
		Left:   Nil,
		Right:  Nil,
	}
}

func (n *Node) UpdateHeight() {
	n.Height = 1 + max(n.Left.Height, n.Right.Height)
	n.BF = n.Right.Height - n.Left.Height
}

func (n *Node) Insert(key int) (*Node, error) {
	var err error
	if n.Key == key {
		err = fmt.Errorf("key %d already exists", key)
	} else if key < n.Key {
		if n.Left.IsNil() {
			n.Left = NewNode(key)
		} else {
			n.Left, err = n.Left.Insert(key)
		}
	} else {
		if n.Right.IsNil() {
			n.Right = NewNode(key)
		} else {
			n.Right, err = n.Right.Insert(key)
		}
	}

	// update height
	n.UpdateHeight()
	log.Printf("node %d bf: %d", n.Key, n.BF)
	n = n.Balance()

	return n, err
}

func (n *Node) RotateRight() *Node {
	x := n.Left
	n.Left = x.Right
	x.Right = n

	n.UpdateHeight()
	x.UpdateHeight()

	return x
}

func (n *Node) RotateLeft() *Node {
	x := n.Right
	n.Right = x.Left
	x.Left = n

	n.UpdateHeight()
	x.UpdateHeight()

	return x
}

func (n *Node) Balance() *Node {
	switch n.BF {
	case -2:
		if n.Left.BF == 1 {
			// LR
			n.Left = n.Left.RotateLeft()
		}
		// LL
		return n.RotateRight()
	case 2:
		if n.Right.BF == -1 {
			// RL
			n.Right = n.Right.RotateRight()
		}
		// RR
		return n.RotateLeft()
	}
	return n
}
