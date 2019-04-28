// crawler project main.go
package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

func main() {
	r := engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	}
	engine.Run(r)
}
