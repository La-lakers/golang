// cdnIp project main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	st := now.Format("2006-01-02")
	fmt.Println(st)

}
