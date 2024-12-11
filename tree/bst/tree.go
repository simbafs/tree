package bst

import (
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

// common methods for Tree

func (t Tree) Root() tree.Node {
	return t.RootNode
}

func (t Tree) Dispatch(cmd string) (tree.Tree, tea.Cmd) {
	seg := strings.Fields(cmd)
	if len(seg) == 0 {
		return t, nil
	}

	switch seg[0] {
	case "insert", "i":
		if len(seg) < 2 {
			return t, tree.ErrMsgf("insert <key>")
		}

		key, err := strconv.Atoi(seg[1])
		if err != nil {
			return t, tree.Cmd(err)
		}

		t.Insert(key)
	}

	return t, nil
}

// custom methods for BST

func (t *Tree) Insert(key int) {
	newNode := NewNode(key)
	if t.RootNode.IsNil() {
		t.RootNode = newNode
	} else {
		t.RootNode.Insert(newNode)
	}
}
