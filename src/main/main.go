// main project main.go
package main

import "fmt"

func main() {
	count := 0
	func(n int) {
		for i := 0; i <= 10; i++ {
			n++
		}
	}(count)
	fmt.Println("count:", count)
	func() {
		for i := 0; i <= 10; i++ {
			count++
		}
	}()
	fmt.Println("count:", count)
}
