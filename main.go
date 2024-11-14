package main

import (
	"fmt"
	"os"

	"tree/tree"
	"tree/tree/bst"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	bstTree := &bst.BST{}
	for _, i := range []int{5, 3, 8, 2, 4, 7, 9, 10000000} {
		bstTree.Insert(i)
	}

	model := tree.New(bstTree)

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
