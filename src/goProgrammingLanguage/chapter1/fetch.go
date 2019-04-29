// fetch.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func fetch(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s:%v", url, err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", b)

}

func main() {
	urls := os.Args[1:]
	if len(urls) == 0 {
		fmt.Fprintln(os.Stderr, "no url")
		os.Exit(1)
	}
	for _, url := range urls {
		fetch(url)
	}
}
