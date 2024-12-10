package aniavl

import (
	"strconv"
	"strings"
	"time"

	"tree/animate"
	"tree/tree"

	tea "github.com/charmbracelet/bubbletea"
)

const cycle = 1000 * time.Millisecond

type Tree struct {
	root *Node

	curr *Node
}

var _ tree.Tree[Node, Tree] = &Tree{}

func (t Tree) Root() *Node {
	return t.root
}

func (t Tree) Dispatch(cmd string) (Tree, tea.Cmd) {
	seg := strings.Fields(cmd)
	if len(seg) == 0 {
		return t, nil
	}

	switch seg[0] {
	case "insert", "i":
		if len(seg) < 2 {
			return t, tree.Msgf("%s <key>", seg[0])
		}

		key, err := strconv.Atoi(seg[1])
		if err != nil {
			return t, tree.Msgf(err.Error())
		}

		t.curr = t.root
		if t.curr != nil {
			t.curr.isCurr = true
		}

		ani := animate.New(cycle)
		ani.Push(BalanceJob(&t.root), InsertJob(key))

		return t, ani.Cmd()
	// case "rotate", "r":
	// 	if len(seg) < 3 {
	// 		return t, tree.Msgf("%s <right|r|left|l> <key>", seg[0])
	// 	}
	//
	// 	if seg[1] != "right" && seg[1] != "r" && seg[1] != "left" && seg[1] != "l" {
	// 		return t, tree.Msgf("invalid direction: %s", seg[1])
	// 	}
	//
	// 	key, err := strconv.Atoi(seg[2])
	// 	if err != nil {
	// 		return t, tree.Msgf(err.Error())
	// 	}
	//
	// 	ani := animate.New(cycle)
	// 	ani.Push(RotateJob{
	// 		key:   key,
	// 		right: seg[1] == "right" || seg[1] == "r",
	// 	})
	//
	// 	return t, ani.Cmd()
	// case "balance", "b":
	// 	if len(seg) < 2 {
	// 		return t, tree.Msgf("insert <key>")
	// 	}
	//
	// 	key, err := strconv.Atoi(seg[1])
	// 	if err != nil {
	// 		return t, tree.Msgf(err.Error())
	// 	}
	//
	// 	ani := animate.New(cycle)
	// 	ani.Push(BalanceJob(key))
	//
	// 	return t, ani.Cmd()
	case "inorder", "in":
		return t, tree.Msgf("%v", t.root.Inorder())
	default:
		return t, tree.Msgf("unknown command: %s", seg[0])

	}
}

func (t Tree) Update(msg tea.Msg) (Tree, tea.Cmd) {
	switch msg := msg.(type) {
	case animate.Msg:
		var cmd tea.Cmd
		var next []tea.Msg
		switch job := msg.Pop().(type) {
		case InsertJob:
			t, next, cmd = insert(t, int(job))
			msg.Push(next...)
			return t, tea.Batch(msg.Cmd(), cmd)
		case RotateJob:
			t, next, cmd = rotate(t, job.node, job.right)
			msg.Push(next...)
			return t, tea.Batch(msg.Cmd(), cmd)
		case BalanceJob:
			t, next, cmd = balance(t, job)
			msg.Push(next...)
			return t, tea.Batch(msg.Cmd(), cmd)
		}
	}

	return t, nil
}

func (t *Tree) SetRoot(root *Node) {
	t.root = root
}
