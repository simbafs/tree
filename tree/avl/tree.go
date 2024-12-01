package avl

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"tree/tree"
)

type Tree struct {
	RootNode *Node
}

var _ tree.Tree[Node] = &Tree{}

func NewTree() *Tree {
	return &Tree{
		RootNode: Nil,
	}
}

// common methods for Tree

func (t *Tree) Root() *Node {
	return t.RootNode
}

func (t *Tree) Op(cmd string) error {
	seg := strings.SplitN(cmd, " ", 2)
	switch seg[0] {
	case "insert", "i":
		if len(seg) < 2 {
			return fmt.Errorf("missing key")
		}

		key, err := strconv.Atoi(seg[1])
		if err != nil {
			return err
		}

		return t.Insert(key)
	case "right-rotate", "r":
		if len(seg) < 2 {
			return fmt.Errorf("missing key")
		}

		node, err := searchNode(t, seg[1])
		if err != nil {
			return err
		}

		t.ReplaceNode(node.Key, node.RotateRight())
	case "left-rotate", "l":
		if len(seg) < 2 {
			return fmt.Errorf("missing key")
		}

		node, err := searchNode(t, seg[1])
		if err != nil {
			return err
		}

		t.ReplaceNode(node.Key, node.RotateLeft())
	case "balance", "b":
		if len(seg) < 2 {
			return fmt.Errorf("missing key")
		}

		node, err := searchNode(t, seg[1])
		if err != nil {
			return err
		}

		t.ReplaceNode(node.Key, node.Balance())
	default:
		return fmt.Errorf("unknown command: %s", cmd)
	}
	return nil
}

// custom methods for Tree

func (t *Tree) Insert(key int) error {
	log.Printf("inserting %d", key)

	if t.RootNode.IsNil() {
		t.RootNode = NewNode(key)
		return nil
	}

	var err error
	t.RootNode, err = t.RootNode.Insert(key)
	return err
}

func (t *Tree) Search(key int) *Node {
	curr := t.Root()

	for !curr.IsNil() {
		if key == curr.Key {
			return curr
		} else if key < curr.Key {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	return nil
}

func (t *Tree) ReplaceNode(key int, node *Node) error {
	curr := t.Root()
	if curr.Key == key {
		t.RootNode = node
		return nil
	}

	for !curr.IsNil() {
		if key == curr.Left.Key {
			curr.Left = node
			return nil
		} else if key == curr.Right.Key {
			curr.Right = node
			return nil
		} else if key < curr.Key {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	return nil
}

// helper functions
func searchNode(t *Tree, str string) (*Node, error) {
	key, err := strconv.Atoi(str)
	if err != nil {
		return nil, err
	}

	node := t.Search(key)
	if node == nil {
		return nil, fmt.Errorf("key %d not found", key)
	}

	return node, nil
}
