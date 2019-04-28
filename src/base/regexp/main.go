// regexp project main.go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := `P[2] condition[all(#1) 7>10] note[DS_5XX报警[lmds]] [tags]keywords:http_status.*?:5.*,filepath:/data/logs/lmds/monitor.log[vals]`

	re, err := regexp.Compile(`not1e\[[^]]+]`)
	if err != nil {
		panic(err)
	}
	result := re.FindString(text)
	if result == "" {
		fmt.Println("nil")
	}
	fmt.Println(result)
}
