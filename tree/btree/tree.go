package btree

import (
	"log"
	"strconv"
	"strings"

	"tree/tree"

	tea "github.com/charmbracelet/bubbletea"
)

type Tree struct {
	root   *Node
	degree int
}

var _ tree.Tree = &Tree{}

func NewTree(degree int) *Tree {
	return &Tree{
		degree: degree,
		root:   NewNode(degree),
	}
}

func init() {
	tree.Register("btree", func() tree.Tree {
		// TODO: degree should be configurable
		return NewTree(2)
	})
}

func (t Tree) Root() tree.Node {
	return t.root
}

func (t Tree) Dispatch(cmd string) (tree.Tree, tea.Cmd) {
	seg := strings.Fields(cmd)
	if len(seg) == 0 {
		return t, nil
	}

	switch seg[0] {
	case "insert", "i":
		if len(seg) < 2 {
			return t, tree.ErrMsgf("%s <key>", seg[0])
		}

		key, err := strconv.Atoi(seg[1])
		if err != nil {
			return t, tree.Cmd(err)
		}

		t.Insert(key)
		return t, nil
	}

	return t, nil
}

func (t Tree) Update(msg tea.Msg) (Tree, tea.Cmd) {
	return t, nil
}

func (t *Tree) SplitRoot() *Node {
	s := NewNode(t.degree)
	// TODO: there will be a bug here
	s.children[0] = t.root
	t.root = s
	s.SplitChild(0)
	return s
}

func (t *Tree) Insert(key int) {
	log.Printf("Insert %d", key)
	r := t.root
	if r.IsFull() {
		r = t.SplitRoot()
	}

	r.InsertNotFull(key)
}
