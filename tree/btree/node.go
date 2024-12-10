package btree

import (
	"fmt"
	"log"

	"tree/tree"
)

type Node struct {
	degree   int
	len      int // the number of keys int the node
	keys     []int
	children []*Node
}

var _ tree.Node[Node] = &Node{}

func NewNode(degree int) *Node {
	return &Node{
		degree:   degree,
		len:      0,
		keys:     make([]int, 2*degree-1),
		children: make([]*Node, 2*degree),
	}
}

func (n Node) View() string {
	s := fmt.Sprintf("%d ", n.keys[:n.len])
	// for i := 0; i < n.len; i++ {
	// 	s += fmt.Sprintf("%d ", n.keys[i])
	// }

	return s
}

func (n Node) Children() []*Node {
	return n.children[:n.len+1]
}

func (n Node) IsNil() bool {
	return false
}

func (n *Node) IsLeaf() bool {
	for _, c := range n.children {
		if c != nil {
			return false
		}
	}
	return true
}

func (n *Node) IsFull() bool {
	return n.len == 2*n.degree-1
}

// Search search if the key in the node.
//   - If found, return the node and true.
//   - If not found, return the node should include the key and false.
func (n *Node) Search(key int) (*Node, bool) {
	i := 0
	for i < n.len && key > n.keys[i] {
		i++
	}

	if i < n.len && key == n.keys[i] {
		return n, true
	} else if n.IsLeaf() {
		return n, false
	} else {
		return n.children[i].Search(key)
	}
}

// SplitChild splits the i-th child of the node.
func (x *Node) SplitChild(i int) {
	y := x.children[i]
	if !y.IsFull() {
		return
	}

	z := NewNode(x.degree)
	z.len = x.degree - 1

	// z gets the last degree-1 keys of y
	for j := 0; j < x.degree-1; j++ {
		z.keys[j] = y.keys[j+x.degree]
	}
	// z gets the last degree children of y
	if !y.IsLeaf() {
		for j := 0; j < x.degree; j++ {
			z.children[j] = y.children[j+x.degree]
		}
	}
	y.len = x.degree - 1

	// shift x's children to the right
	for j := x.len; j >= i+1; j-- {
		x.children[j+1] = x.children[j]
	}
	x.children[i+1] = z

	// shift x's keys to the right
	for j := x.len - 1; j >= i; j-- {
		x.keys[j+1] = x.keys[j]
	}
	x.keys[i] = y.keys[x.degree-1] // insert y's median key to the i-th key of x

	x.len++
}

func (x *Node) InsertNotFull(key int) {
	log.Printf("insert %d on %v", key, x.keys)
	i := x.len - 1
	if x.IsLeaf() {
		// move keys to the right to make space for the new key
		for i >= 0 && key < x.keys[i] {
			x.keys[i+1] = x.keys[i]
			i--
		}
		x.keys[i+1] = key
		x.len++
	} else {
		// find the child to insert the key
		for i >= 0 && key < x.keys[i] {
			i--
		}
		i++
		if x.children[i].IsFull() {
			x.SplitChild(i)
			if key > x.keys[i] {
				i++
			}
		}
		x.children[i].InsertNotFull(key)
	}
}
