package main

import (
	"fmt"
	"time"
)

// Case B:
// https://go101.org/article/channel.html

func main () {
	c := make(chan int) // an unbuffered channel

	go func(ch chan<- int, x int) {
		time.Sleep(time.Second)
		// <-ch  // fails to compile
		// Send the value and block until the result is received.
		ch <- x * x // 9 is sent
		//fmt.Println("time")
	}(c, 3)
	//---//----//
	done := make(chan struct{})
	go func(ch <-chan int) {
		n := <-ch
		fmt.Println(n)

		time.Sleep(time.Second)
		done <- struct{}{}
	}(c)

	<-done
	fmt.Println("bye")
}
