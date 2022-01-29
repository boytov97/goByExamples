package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func multipleReturns(a int, str string) (int, string) {
	return a, str
}

func multipleReturns2() (int, string, string, bool) {
	return 14, "hello", "world", false
}

func main () {

	a, b := vals()
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	_, c := vals()
	fmt.Println("c:", c)

	numBer, strIng := multipleReturns(a, "string")
	fmt.Println("num:", numBer)
	fmt.Println("str:", strIng)

	count, word1, word2, _ := multipleReturns2()
	fmt.Printf("count: %d, word1: %s, word2: %s \n", count, word1, word2)

}
