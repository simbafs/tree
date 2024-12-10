package bst

import (
	"fmt"
	"strconv"
	"strings"

	"tree/tree"

	tea "github.com/charmbracelet/bubbletea"
)

type Tree struct {
	RootNode *Node
}

var _ tree.Tree[Node, Tree] = &Tree{}

// common methods for Tree

func (t Tree) Root() *Node {
	return t.RootNode
}

func (t Tree) Dispatch(cmd string) (Tree, tea.Cmd) {
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
			return t, tree.ErrMsg(err)
		}

		t.Insert(key)
	}

	return t, nil
}

func (t Tree) Update(msg tea.Msg) (Tree, tea.Cmd) {
	return t, nil
}

// custom methods for BST

func (t *Tree) Insert(key int) {
	newNode := &Node{Key: key}
	if t.RootNode == nil {
		t.RootNode = newNode
	} else {
		t.RootNode.Insert(newNode)
	}
}

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

var _ tree.Node[Node] = &Node{}

// common methods for Node

func (node Node) View() string {
	return fmt.Sprintf("%d", node.Key)
}

func (node Node) Children() []*Node {
	return []*Node{node.Left, node.Right}
}

func (node Node) IsNil() bool {
	return false
}

// custom methods for BSTNode

func (node *Node) Insert(newNode *Node) {
	if newNode.Key < node.Key {
		if node.Left == nil {
			node.Left = newNode
		} else {
			node.Left.Insert(newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			node.Right.Insert(newNode)
		}
	}
}
