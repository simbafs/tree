package bst

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type BST struct {
	RootNode *BSTNode
}

// common methods for Tree

func (t BST) Root() *BSTNode {
	return t.RootNode
}

func (t *BST) Op(cmd string) error {
	seg := strings.Fields(cmd)
	if len(seg) == 0 {
		return nil
	}

	switch seg[0] {
	case "insert", "i":
		if len(seg) < 2 {
			return fmt.Errorf("insert <key>")
		}

		key, err := strconv.Atoi(seg[1])
		if err != nil {
			return err
		}

		t.Insert(key)
	}

	return nil
}

// custom methods for BST

func (t *BST) Insert(key int) {
	newNode := &BSTNode{Key: key}
	if t.RootNode == nil {
		t.RootNode = newNode
	} else {
		t.RootNode.Insert(newNode)
	}
}

type BSTNode struct {
	Key   int
	Left  *BSTNode
	Right *BSTNode
}

// common methods for Node

func (node BSTNode) View() string {
	return fmt.Sprintf("%d", node.Key)
}

func (node BSTNode) Children() []*BSTNode {
	return []*BSTNode{node.Left, node.Right}
}

func (node BSTNode) Styling(s lipgloss.Style) lipgloss.Style {
	return s
}

// custom methods for BSTNode

func (node *BSTNode) Insert(newNode *BSTNode) {
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
