// done project main.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func chanDemo() {
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(id int) {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			fmt.Printf("worker 0 recevied %d\n", id)

			done <- true
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

func main() {
	cha := make(chan<- int)
	go func() {
		for {
			cha <- 1
		}
	}()
	time.Sleep(time.Second)
	//chanDemo()

}
