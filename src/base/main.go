// base project main.go
package main

import (
	"flag"
	"fmt"
)

func main() {
	var filenames string
	flag.StringVar(&filenames, "f", "", "filenames")
	flag.Parse()
	fmt.Print(filenames)

}
