package main

import (
	"fmt"
	"os"

	"tree/tree"
	"tree/tree/avl"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	tea.LogToFile("log.txt", "")

	avlTree := avl.NewTree()
	// bstTree := &bst.Tree{}
	for _, i := range []int{4, 5, 6, 1, 2} {
		avlTree.Insert(i)
		// bstTree.Insert(i)
	}

	model := tree.New(avlTree)
	// model := tree.New(bstTree)

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
