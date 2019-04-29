// ex1.4.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("no file")
		return
	}
	for _, arg := range args {
		counts := make(map[string]int)
		func(file string, counts map[string]int) {
			f, err := os.Open(file)
			if err != nil {
				fmt.Printf("read  %s error: %v", file, err)
				return
			}
			defer f.Close()
			input := bufio.NewScanner(f)
			for input.Scan() {
				if _, ok := counts[input.Text()]; ok {
					fmt.Println(file)
					return
				} else {
					counts[input.Text()]++
				}
			}
		}(arg, counts)

	}
}
