package main

import "fmt"

func iniSeq() func() int {
	i := 0

	return func() int {
		i++

		return i
	}
}

func main() {
	nextInt := iniSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := iniSeq()
	fmt.Println(newInts())
}
