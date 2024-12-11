package main

import (
	"fmt"
	"os"

	"tree/tree"
	"tree/tree/bst"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	tea.LogToFile("log.txt", "")

	// avlTree := avl.NewTree()
	bstTree := bst.NewTree()
	// bTree := btree.NewTree(2)

	// bTree := btree.NewTree(2)
	// for _, i := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12} {
	// 	avlTree.Insert(i)
	// 	// bstTree.Insert(i)
	// 	// bTree.Insert(i)
	// }

	// model := tree.New(avlTree)
	model := tree.New(bstTree)
	// model := tree.New(bTree)

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
