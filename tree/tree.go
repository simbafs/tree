package tree

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Define Tree and Node interfaces as before.
type Tree[T Node[T]] interface {
	Root() *T
	Op(cmd string) error // pass a command to operate on the tree
}

type Node[T any] interface {
	View() string // print the node without border
	Children() []*T
}

// Define ModelTree as a generic struct.
type ModelTree[T Node[T]] struct {
	tree Tree[T]
	cmd  textinput.Model
	msg  string
}

func New[T Node[T]](tree Tree[T]) ModelTree[T] {
	cmd := textinput.New()
	cmd.Focus()

	return ModelTree[T]{
		tree: tree,
		cmd:  cmd,
	}
}

// Implement Init method for ModelTree.
func (t ModelTree[T]) Init() tea.Cmd {
	return nil
}

// Implement Update method for ModelTree.
func (t ModelTree[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return t, tea.Quit
		case "enter":
			switch t.cmd.Value() {
			case "quit", "exit":
				return t, tea.Quit
			default:
				if err := t.tree.Op(t.cmd.Value()); err != nil {
					t.msg = err.Error()
				}
			}
			t.cmd.Reset()

		case "esc":
			t.cmd.Reset()
		}
	}

	var cmd tea.Cmd
	t.cmd, cmd = t.cmd.Update(msg)

	return t, cmd
}

// Implement View method for ModelTree.
func (t ModelTree[T]) View() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		ModelNode[T]{t.tree.Root()}.View(),
		t.cmd.View(),
		t.msg,
		"",
	)
}

// Define ModelNode as a generic struct to view a single node.
type ModelNode[T Node[T]] struct {
	node *T
}

// Implement the View method for ModelNode.
func (node ModelNode[T]) View() string {
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Align(lipgloss.Center)

	if node.node == nil {
		return style.Render("")
	}

	n := *node.node

	atLeastOneChild := false
	children := []string{}
	widthOfChildren := 0

	for _, child := range n.Children() {
		if child != nil {
			atLeastOneChild = true
		}
		c := ModelNode[T]{child}.View()
		children = append(children, c)
		widthOfChildren += lipgloss.Width(c)
	}

	self := style.Width(max(widthOfChildren-2, lipgloss.Width(n.View()))). // -2 for the border
										Render(n.View())

	if atLeastOneChild {
		return lipgloss.JoinVertical(lipgloss.Center,
			self,
			lipgloss.JoinHorizontal(lipgloss.Top, children...),
		)
	} else {
		return self
	}
}
