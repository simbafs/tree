package tree

import tea "github.com/charmbracelet/bubbletea"

type T interface {
	Root() any
	Dispatch(string) (T, tea.Cmd)
	Update(tea.Msg) (T, tea.Cmd)
}

var trees = map[string]func() *T{}

func Register(name string, tree func() *T) {
	trees[name] = tree
}

func GetTree(name string) (*T, bool) {
	if t, ok := trees[name]; ok {
		return t(), true
	}
	return nil, false
}
