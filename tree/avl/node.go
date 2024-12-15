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

	isNil bool
}

var _ tree.Node = &Node{}

// common methods for Node

func (node Node) View() string {
	return fmt.Sprintf("%d (h: %d, %d)", node.Key, node.Height, node.BF)
}

func (node Node) Children() []tree.Node {
	return []tree.Node{node.Left, node.Right}
}

func (node Node) IsNil() bool {
	return node.isNil
}

// custom methods for Node

var Nil = &Node{
	Key:    -1,
	Height: 0,
	isNil:  true,
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

func (n *Node) Delete(key int) (*Node, error) {
	var err error
	if n.IsNil() {
		err = fmt.Errorf("key %d not found", key)
	} else if key < n.Key {
		n.Left, err = n.Left.Delete(key)
	} else if key > n.Key {
		n.Right, err = n.Right.Delete(key)
	} else {
		if n.Left.IsNil() {
			return n.Right, nil
		} else if n.Right.IsNil() {
			return n.Left, nil
		}

		// find the smallest node in the right subtree
		min := n.Right
		for !min.Left.IsNil() {
			min = min.Left
		}

		// replace the current node with the smallest node
		n.Key = min.Key
		n.Right, err = n.Right.Delete(min.Key)
	}

	n.UpdateHeight()
	n = n.Balance()

	return n, err
}
