package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, seq := "", ""
	for _, i := range os.Args[1:] {
		s += seq + i
		seq = "===="
	}
	fmt.Println(s)
	fmt.Println(strings.Join(os.Args[1:], ","))
}
