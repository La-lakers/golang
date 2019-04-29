// ex1.7.go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetch(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("fetch: %v", err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
func main() {
	args := os.Args[1:]
	for _, url := range args {
		fetch(url)
	}
}
