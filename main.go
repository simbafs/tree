package main

import (
	"fmt"
	"log"
	"os"

	"tree/tree"
	"tree/tree/aniavl"

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
	avlTree := &aniavl.Tree{}

	root := aniavl.NewNode(5)
	root.SetLeft(aniavl.NewNode(4))
	root.SetRight(aniavl.NewNode(6))
	avlTree.SetRoot(root)

	// for _, v := range []int{5, 3, 7, 2, 4, 6, 8} {
	// 	avlTree.Dispatch(fmt.Sprintf("insert %d", v))
	// }

	model := tree.New(avlTree)

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
