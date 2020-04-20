package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go looper(ch)
	select {
	case res := <-ch:
		fmt.Println(res)

		//default:

	}

}

func looper(c chan int) {
	for i := 0; i < 100; i++ {
		time.Sleep(2 * time.Second)
		c <- i
	}
}
