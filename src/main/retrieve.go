// retrieve
package main

import (
	"fake"
	"fmt"
)

const url = "http://www.baidu.com"

type Retrieve interface {
	Get(url string) string
}

func download(r Retrieve) string {
	return r.Get(url)
}
func inspect(r Retrieve) string{
	switch r.(type)
	case fake.Retrieve:
	
}

func main() {
	var r Retrieve
	r = fake.Retrieve{}
	fmt.Println(download(r))
}
