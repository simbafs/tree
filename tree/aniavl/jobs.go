package aniavl

import (
	"tree/tree"

	tea "github.com/charmbracelet/bubbletea"
)

// Jobs
type InsertJob int

func insert(t Tree, key int) (Tree, []tea.Msg, tea.Cmd) {
	if t.root == nil {
		t.root = NewNode(key)
		return t, nil, tree.Msgf("insert %d as root", key)
	}

	t.curr.isCurr = false

	if t.curr.value == key {
		return t, nil, tree.Msgf("key %d already exists", key)
	} else if key < t.curr.value {
		if t.curr.left.IsNil() {
			t.curr.left = NewNode(key)
			return t, []tea.Msg{BalanceJob(t.curr.left)}, tree.Msgf("insert %d as left child of %d", key, t.curr.value)
		}

		t.curr = t.curr.left
		t.curr.isCurr = true
		return t, []tea.Msg{BalanceJob(t.curr), InsertJob(key)}, tree.Msgf("travel to left child of %d", t.curr.value)
	} else {
		if t.curr.right.IsNil() {
			t.curr.right = NewNode(key)
			return t, []tea.Msg{BalanceJob(t.curr.right)}, tree.Msgf("insert %d as right child of %d", key, t.curr.value)
		}

		t.curr = t.curr.right
		t.curr.isCurr = true

		return t, []tea.Msg{BalanceJob(t.curr), InsertJob(key)}, tree.Msgf("travel to right child of %d", t.curr.value)
	}
}

type RotateJob struct {
	node  *Node
	right bool
}

func rotate(t Tree, node *Node, right bool) (Tree, []tea.Msg, tea.Cmd) {
	return t, nil, nil
}

func rotateRight(node *Node) *Node {
	return nil
}

func rotateLeft(node *Node) *Node {
	return nil
}

type BalanceJob *Node

func balance(t Tree, node *Node) (Tree, []tea.Msg, tea.Cmd) {
	if node == nil {
		return t, nil, nil
	}

	node.UpdateHeight()

	next := []tea.Msg{}

	if node.bf > 1 {
		if node.right.bf < 0 {
			next = append(next, RotateJob{node.right, true})
		}
		next = append(next, RotateJob{node, false})
	} else if node.bf < -1 {
		if node.left.bf > 0 {
			next = append(next, RotateJob{node.left, false})
		}
		next = append(next, RotateJob{node, true})
	}

	return t, next, tree.Msgf("balance %d", node.value)
}
