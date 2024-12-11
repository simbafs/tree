package main

import (
	"fmt"
	"os"
	"strings"

	"tree/tree"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	_ "tree/tree/avl"
	_ "tree/tree/bst"
	_ "tree/tree/btree"
)

type Model struct {
	all   []string
	input textinput.Model
	msg   string
}

func NewModel() Model {
	text := textinput.New()
	text.Focus()
	text.Placeholder = "Enter tree model"

	return Model{
		all:   tree.AllTrees(),
		input: text,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			tree, ok := tree.GetTreeModel(m.input.Value())
			if ok {
				return tree, nil
			} else {
				m.msg = "Unknown tree model"
			}

		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		strings.Join(m.all, " / "),
		m.input.View(),
		m.msg,
	)
}

func main() {
	tea.LogToFile("log.txt", "")

	model := NewModel()

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
