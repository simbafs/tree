# Tree

Tree is a project to visuiallize tree like data structure

# Supported tree

-   BST
-   AVL
-   B-Tree

# The Problem of Null Pointer

In golang, it is expected that the value assigned to an interface implement the methods, so `nil` should never be assigned to an interface. To repersent a null pointer(e.g. a node has no child), every pointer should point to some node, either data node or a `Nil` Node.

For example, `Node` is defined like this

```go
type Node struct {
    // some fields
    left *Node
    right *Node

    isNil bool
}

// some methods

var Nil = &Node{
    isNil: True
}

func NewNode() *Node {
    return &Node{
        left: Nil,
        right: Nil,
    }
}

func (n Node) IsNil() bool {
    return n.isNil
}
```

The pointer `left` and `right` of every normal node point to Nil. So you can and should always check if a node is null with `node.IsNil()`, not `node == nil`

# TODO

-   [x] redesign IsNil
-   [ ] Animation
-   [ ] An UI to select which registered tree to use
