package anibst

import (
	"fmt"

	"tree/tree"

	"github.com/charmbracelet/lipgloss"
)

type Node struct {
	value int
	left  *Node
	right *Node

	isCurr bool
}

// check if Node implements tree.Node
var _ tree.Node[Node] = &Node{}

func (n Node) View() string {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff"))

	if n.isCurr {
		style = style.Background(lipgloss.Color("#00aa00"))
	}

	return style.Render(fmt.Sprintf("%d", n.value))
}

func (n Node) Children() []*Node {
	return []*Node{n.left, n.right}
}

func (n Node) IsNil() bool {
	return false
}
