package main

import "fmt"

func main() {
	c := make(chan int, 2) // a buffered channel
	c <- 3
	c <- 5

	close(c)
	fmt.Println(len(c), cap(c)) // 2 2

	x, ok := <-c
	fmt.Println(x, ok)          // 3 true
	fmt.Println(len(c), cap(c)) // 1 2

	x, ok = <-c
	fmt.Println(x, ok)          // 5 true
	fmt.Println(len(c), cap(c)) // 0 2

	//close(c)

	//c <- 7
}
