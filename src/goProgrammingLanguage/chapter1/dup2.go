// dup2.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func main() {
	counts := make(map[string]int)
	args := os.Args[1:]
	if len(args) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range args {
			func(file string) {
				f, err := os.Open(file)
				if err != nil {
					fmt.Printf("read %s error:%v", file, err)
					return
				}
				defer f.Close()
				countLines(f, counts)
			}(arg)
		}
	}
	for k, v := range counts {
		fmt.Printf("line:%s\t count:%d\n", k, v)
	}

}
