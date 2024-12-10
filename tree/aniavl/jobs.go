package aniavl

import (
	"log"

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
			return t, []tea.Msg{BalanceJob(&t.curr.left)}, tree.Msgf("insert %d as left child of %d", key, t.curr.value)
		}

		t.curr = t.curr.left
		t.curr.isCurr = true
		return t, []tea.Msg{BalanceJob(&t.curr), InsertJob(key)}, tree.Msgf("travel to left child of %d", t.curr.value)
	} else {
		if t.curr.right.IsNil() {
			t.curr.right = NewNode(key)
			return t, []tea.Msg{BalanceJob(&t.curr.right)}, tree.Msgf("insert %d as right child of %d", key, t.curr.value)
		}

		t.curr = t.curr.right
		t.curr.isCurr = true

		return t, []tea.Msg{BalanceJob(&t.curr), InsertJob(key)}, tree.Msgf("travel to right child of %d", t.curr.value)
	}
}

type RotateJob struct {
	node  **Node
	right bool
}

func rotate(t Tree, node **Node, right bool) (Tree, []tea.Msg, tea.Cmd) {
	if *node == nil {
		return t, nil, nil
	}

	log.Printf("before rotate, node is %d", (*node).value)
	if right {
		*node = rotateRight(*node)
	} else {
		*node = rotateLeft(*node)
	}
	log.Printf("after rotate, node is %d", (*node).value)

	return t, nil, nil
}

func rotateRight(node *Node) *Node {
	log.Printf("rotating right %d", node.value)
	x := node.left
	node.left = x.right
	x.right = node

	node.UpdateHeight()
	x.UpdateHeight()

	return x
}

func rotateLeft(node *Node) *Node {
	log.Printf("rotating left %d", node.value)
	x := node.right
	node.right = x.left
	x.left = node

	node.UpdateHeight()
	x.UpdateHeight()

	return x
}

type BalanceJob **Node

func balance(t Tree, node **Node) (Tree, []tea.Msg, tea.Cmd) {
	if (*node) == nil {
		return t, nil, nil
	}

	log.Printf("balancing %d", (*node).value)

	n := *node

	n.UpdateHeight()

	if n.bf > 1 {
		if n.right.bf < 0 {
			return t, []tea.Msg{RotateJob{node, false}, RotateJob{&n.right, true}}, tree.Msgf("balance %d (RL)", n.value)
		} else {
			return t, []tea.Msg{RotateJob{node, false}}, tree.Msgf("balance %d (RR)", n.value)
		}
	} else if n.bf < -1 {
		if n.left.bf > 0 {
			return t, []tea.Msg{RotateJob{node, true}, RotateJob{&n.left, false}}, tree.Msgf("balance %d (LR)", n.value)
		} else {
			return t, []tea.Msg{RotateJob{node, true}}, tree.Msgf("balance %d (LL)", n.value)
		}
	}

	return t, nil, tree.Msgf("balance %d", n.value)
}
