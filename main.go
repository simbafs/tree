package main

import (
	"fmt"
	"log"
	"os"

	"tree/tree"
	"tree/tree/bst"

	tea "github.com/charmbracelet/bubbletea"
)

func init() {
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY|os.O_SYNC|os.O_CREATE, 0o664)
	if err != nil {
		panic(err)
	}

	log.SetOutput(f)

	log.Println("starting...")
}

func main() {
	// avlTree := avl.NewTree()
	bstTree := &bst.BST{}
	for _, i := range []int{4, 5, 6, 1, 2} {
		// avlTree.Insert(i)
		bstTree.Insert(i)
	}

	// model := tree.New(avlTree)
	model := tree.New(bstTree)

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
