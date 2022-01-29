package main

import "fmt"

func printHello () {
	fmt.Println("Hello from printHello")
}

func main () {
	go func() {fmt.Println("Hello Inline")}()

	go printHello()

	fmt.Println("Hello from main")
}
