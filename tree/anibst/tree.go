package anibst

import (
	"log"
	"strconv"
	"strings"
	"time"

	"tree/tree"

	tea "github.com/charmbracelet/bubbletea"
)

const cycle = 1 * time.Second

type Tree struct {
	root *Node

	curr *Node
}

// check if Tree implements tree.Tree
var _ tree.Tree[Node, Tree] = &Tree{}

func (t Tree) Root() *Node {
	return t.root
}

func (t Tree) Dispatch(cmd string) (Tree, tea.Cmd) {
	seg := strings.Fields(cmd)
	if len(seg) == 0 {
		return t, tree.Msgf("empty command")
	}

	switch seg[0] {
	case "insert", "i":
		if len(seg) < 2 {
			return t, tree.Msgf("insert <key>")
		}

		key, err := strconv.Atoi(seg[1])
		if err != nil {
			return t, tree.Msgf(err.Error())
		}

		log.Printf("insert %d", key)

		if t.root == nil {
			t.root = &Node{value: key}
			return t, tree.Msgf("insert %d", key)
		} else {
			t.curr = t.root
			t.curr.isCurr = true
			return t, tea.Batch(Insert(key), tree.Msgf("search root"))
		}
	default:
		return t, tree.Msgf("unknown command")
	}
}

func (t Tree) Update(msg tea.Msg) (Tree, tea.Cmd) {
	switch msg := msg.(type) {
	case InsertJob:
		log.Printf("insert %d on %d", msg.value, t.curr.value)
		t.curr.isCurr = false
		if msg.value < t.curr.value {
			if t.curr.left == nil {
				t.curr.left = &Node{value: msg.value}
				return t, tree.Msgf("insert %d", msg.value)
			} else {
				v := t.curr.value
				t.curr = t.curr.left
				t.curr.isCurr = true
				return t, tea.Batch(Insert(msg.value), tree.Msgf("search %d's left", v))
			}
		} else {
			if t.curr.right == nil {
				t.curr.right = &Node{value: msg.value}
				return t, tree.Msgf("insert %d", msg.value)
			} else {
				v := t.curr.value
				t.curr = t.curr.right
				t.curr.isCurr = true
				return t, tea.Batch(Insert(msg.value), tree.Msgf("search %d's right", v))
			}
		}
	}

	return t, nil
}

type InsertJob struct {
	value int
}

func Insert(value int) tea.Cmd {
	return tea.Tick(cycle, func(time.Time) tea.Msg {
		return InsertJob{value}
	})
}
