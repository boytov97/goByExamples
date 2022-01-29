package main

import "fmt"

func main () {
	m := make(map[string]int)
	fmt.Println("map:", m)

	m["k1"] = 7
	m["k2"] = 13
	m["k3"] = 15

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	//_, prs := m["k2"]
	//fmt.Println("")
}
