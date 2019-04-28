// echo1.go
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	var s, seq string
	for i := 0; i < len(args); i++ {
		s += seq + args[i]
		seq = ","

	}
	fmt.Println(s)
}
