package main

import (
	"fmt"
	"time"
)

func main() {
	// make a channel
	ch := make(chan int)
	defer close(ch)

	go looper(ch)
	for {
		select {
		case res := <-ch:
			fmt.Println(" ", res)

		default:

		}
	}

}

func looper(c chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		c <- i
	}
}
