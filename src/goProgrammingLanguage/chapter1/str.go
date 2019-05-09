// str.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = []string{"有影响",
		"影响业务",
		"有问题",
		"有异常",
		"有质量影响",
		"影响卡顿"}
	s1 := "24号报警。404报警多数是大主播下线造成，和200比例报警也在优化，有影响"
	for _, s2 := range s {
		fmt.Printf("%s\t%t\n", s2, strings.Contains(s1, s2))
	}
}
