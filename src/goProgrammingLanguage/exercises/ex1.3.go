// ex1.3.go
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func join() {
	fmt.Println(strings.Join(os.Args, ","))
}

func xunhuan() {
	s, seq := "", ""
	for _, i := range os.Args {
		s += seq + i
		seq = ","
	}
	fmt.Println(s)

}
func main() {
	start := time.Now().UnixNano()
	fmt.Println(start)
	join()
	end := time.Now().UnixNano()
	fmt.Println(end)
	fmt.Println(end - start)
}
