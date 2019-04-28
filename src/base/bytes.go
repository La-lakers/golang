// bytes.go
package main

import (
	"fmt"
)

func main() {
	s := []byte("a b c d ")
	fmt.Printf("%v\n", s)
	s2 := []byte("A B C D")
	s = append(s, s2...)

	fmt.Printf("%v\n", s2)
	fmt.Printf("%s\n", s)
}
