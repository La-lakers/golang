// tree project tree.go
package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (n *Node) Print() {
	fmt.Print(n.Value, " ")
}

func (n *Node) SetValue(value int) {
	n.Value = value
}

func (n *Node) Traverse() {
	if n == nil {
		return
	}
	n.Left.Traverse()
	n.Print()
	n.Right.Traverse()
}

func (n *Node) TraverseFunc(f func(*Node)) {
	if n == nil {
		return
	}
	n.Left.TraverseFunc(f)
	f(n)
	n.Right.TraverseFunc(f)
}
