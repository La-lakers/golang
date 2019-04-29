// cup3.go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func countLines(b []byte, counts map[string]int) {
	b1 := strings.Split(string(b), "\n")
	for _, line := range b1 {
		//line = strings.Trim(line, "")
		counts[line]++
	}
}

func main() {
	args := os.Args[1:]
	counts := make(map[string]int)
	if len(args) == 0 {
		fmt.Println("no file")
		return
	}
	for _, file := range args {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("read %s error:%v", file, err)
			continue
		}
		countLines(b, counts)
	}
	for k, v := range counts {
		fmt.Printf("Line:%s\tCount:%d\n", k, v)
	}
}
