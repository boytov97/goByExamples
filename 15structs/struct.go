package main

import "fmt"

type person struct {
	name string
	age int
}

func NewPerson(name string) *person {

	p := person{name: name}
	p.age = 46

	return &p
}

func main () {

	fmt.Println(person{name: "Bob"})

	fmt.Println(person{name: "Alice", age: 21})

	fmt.Println(person{name: "Fred"})

	fmt.Println(person{name: "Ann", age: 24})

	fmt.Println(NewPerson("Jon"))

	s := person{name: "Asan", age: 45}
	fmt.Println("name:", s.name)

	sp := &s
	fmt.Println("Asan's age:", sp.age)

	sp.age = 35
	fmt.Println("Asan's age:", s.age)
}
