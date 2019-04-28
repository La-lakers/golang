// ex1.1.go
package main

import (
	"fmt"
	"os"
)

func main() {
	s, seq := "", ""
	for _, arg := range os.Args {
		s += seq + arg
		seq = " "
	}
	fmt.Println(s)
}
