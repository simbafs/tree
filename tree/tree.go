package tree

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Define Tree and Node interfaces as before.
type Tree interface {
	Root() Node
	Dispatch(string) (Tree, tea.Cmd)
}

type Node interface {
	View() string // print the node without border
	Children() []Node
	IsNil() bool
}

// Define ModelTree as a generic struct.
type ModelTree struct {
	tree Tree
	cmd  textinput.Model
	msg  string
}

func New(tree Tree) *ModelTree {
	cmd := textinput.New()
	cmd.Focus()

	return &ModelTree{
		tree: tree,
		cmd:  cmd,
	}
}

// Implement Init method for ModelTree.
func (t ModelTree) Init() tea.Cmd {
	return nil
}

// Implement Update method for ModelTree.
func (t ModelTree) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case Msg:
		t.msg = string(msg)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return t, tea.Quit
		case "enter":
			switch t.cmd.Value() {
			case "quit", "exit":
				return t, tea.Quit
			default:
				t.tree, cmd = t.tree.Dispatch(t.cmd.Value())
				cmds = append(cmds, cmd)
			}
			t.cmd.Reset()

		case "esc":
			t.cmd.Reset()
		}
	}

	t.cmd, cmd = t.cmd.Update(msg)
	cmds = append(cmds, cmd)

	return t, tea.Batch(cmds...)
}

// Implement View method for ModelTree.
func (t ModelTree) View() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		ModelNode{t.tree.Root()}.View(),
		t.cmd.View(),
		t.msg,
		"",
	)
}

// Define ModelNode as a generic struct to view a single node.
type ModelNode struct {
	node Node
}

// Implement the View method for ModelNode.
func (node ModelNode) View() string {
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Align(lipgloss.Center)

	if node.node.IsNil() {
		return style.Render("")
	}

	n := node.node

	atLeastOneChild := false
	children := []string{}
	widthOfChildren := 0

	for _, child := range n.Children() {
		if !child.IsNil() {
			atLeastOneChild = true
		}
		c := ModelNode{child}.View()
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
	}

	return self
}
