package aniavl

import (
	"strconv"
	"strings"

	"tree/tree"

	tea "github.com/charmbracelet/bubbletea"
)

type Tree struct {
	root *Node
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

		return t, Job(InsertJob(key))
	case "rotate", "r":
		if len(seg) < 3 {
			return t, tree.Msgf("%s <right|r|left|l> <key>", seg[0])
		}

		if seg[1] != "right" && seg[1] != "r" && seg[1] != "left" && seg[1] != "l" {
			return t, tree.Msgf("invalid direction: %s", seg[1])
		}

		key, err := strconv.Atoi(seg[2])
		if err != nil {
			return t, tree.Msgf(err.Error())
		}

		return t, Job(RotateJob{
			key:   key,
			right: seg[1] == "right" || seg[1] == "r",
		})
	case "balance", "b":
		if len(seg) < 2 {
			return t, tree.Msgf("insert <key>")
		}

		key, err := strconv.Atoi(seg[1])
		if err != nil {
			return t, tree.Msgf(err.Error())
		}

		return t, Job(BalanceJob(key))
	default:
		return t, tree.Msgf("unknown command: %s", seg[0])

	}
}

func (t Tree) Update(msg tea.Msg) (Tree, tea.Cmd) {
	switch msg := msg.(type) {
	case InsertJob:

	case RotateJob:

	case BalanceJob:

	}

	return t, nil
}

// Jobs

func Job[T any](job T) tea.Cmd {
	return func() tea.Msg {
		return job
	}
}

type InsertJob int

type RotateJob struct {
	key   int
	right bool
}

type BalanceJob int
