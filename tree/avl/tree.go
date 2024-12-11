package avl

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"tree/tree"

	tea "github.com/charmbracelet/bubbletea"
)

type Tree struct {
	RootNode *Node
}

var _ tree.Tree = &Tree{}

func NewTree() *Tree {
	return &Tree{
		RootNode: Nil,
	}
}

func init() {
	// tree.Register("avl", NewTree)
}

// common methods for Tree

func (t Tree) Root() tree.Node {
	return t.RootNode
}

func (t Tree) Dispatch(cmd string) (tree.Tree, tea.Cmd) {
	seg := strings.SplitN(cmd, " ", 2)
	switch seg[0] {
	case "insert", "i":
		if len(seg) < 2 {
			return t, tree.ErrMsgf("missing key")
		}

		key, err := strconv.Atoi(seg[1])
		if err != nil {
			return t, tree.Cmd(err)
		}

		return t, tree.Cmd(t.Insert(key))
	case "right-rotate", "r":
		if len(seg) < 2 {
			return t, tree.ErrMsgf("missing key")
		}

		node, err := searchNode(&t, seg[1])
		if err != nil {
			return t, tree.Cmd(err)
		}

		t.ReplaceNode(node.Key, node.RotateRight())
	case "left-rotate", "l":
		if len(seg) < 2 {
			return t, tree.ErrMsgf("missing key")
		}

		node, err := searchNode(&t, seg[1])
		if err != nil {
			return t, tree.Cmd(err)
		}

		t.ReplaceNode(node.Key, node.RotateLeft())
	case "balance", "b":
		if len(seg) < 2 {
			return t, tree.ErrMsgf("missing key")
		}

		node, err := searchNode(&t, seg[1])
		if err != nil {
			return t, tree.Cmd(err)
		}

		t.ReplaceNode(node.Key, node.Balance())
	default:
		return t, tree.ErrMsgf("unknown command: %s", cmd)
	}
	return t, nil
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
	curr := t.RootNode

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
	curr := t.RootNode
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
	if node.IsNil() {
		return nil, fmt.Errorf("key %d not found", key)
	}

	return node, nil
}
