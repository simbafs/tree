package tree

var trees = map[string]func() Tree{}

func Register(name string, tree func() Tree) {
	trees[name] = tree
}

func GetTreeModel(name string) (*ModelTree, bool) {
	if t, ok := trees[name]; ok {
		return New(t()), ok
	}
	return nil, false
}

func AllTrees() []string {
	var names []string
	for name := range trees {
		names = append(names, name)
	}
	return names
}
