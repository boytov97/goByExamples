package main

import "fmt"

func zeroval (ival int) {
	ival = 0

	fmt.Println("ival in:", ival)
}

func zeroptr (iptr *int) {
	*iptr = 0
}

func main () {

	i := 1
	zeroval(i)
	fmt.Println("i val:", i)

	zeroptr(&i)
	fmt.Println("i ptr:", i)

	fmt.Println("pointer:", &i)
}
