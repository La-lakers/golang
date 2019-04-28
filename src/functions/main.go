// functions project main.go
package main

import (
	"fmt"
	"tree"
)

func addFunc() func(n int) int {
	sum := 0
	return func(n int) int {
		sum += n
		return sum
	}
}

func main() {
	f := addFunc()
	for i := 1; i <= 10; i++ {
		fmt.Printf("0 + 1 + ...%d = %d\n", i, f(i))
	}
	t := &tree.Node{10, nil, nil}
	t.Left = &tree.Node{5, nil, nil}
	t.Right = &tree.Node{15, nil, nil}
	t.Left.Right = &tree.Node{6, nil, nil}
	t.Right.Left = &tree.Node{13, nil, nil}
	t.Right.Left.Right = &tree.Node{14, nil, nil}
	fmt.Println(t)
	t.SetValue(11)
	fmt.Println(t)4
	t.Traverse()
	fmt.Println()
	countValue := 0
	t.TraverseFunc(func(t *tree.Node) {
		t.Print()
		countValue++
	})
	fmt.Println()
	fmt.Println("CountValue:", countValue)

}
