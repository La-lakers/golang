// ex1.8.go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func fetch(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	fmt.Println("Response Status:", resp.StatusCode) //ex1.9
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "no url\n")
		return
	}
	for _, url := range args {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		fetch(url)
	}
}
