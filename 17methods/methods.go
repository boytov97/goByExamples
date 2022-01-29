package main

import "fmt"

type rect struct {
	width, height int
}

// Метод area принимает получателя *rect.
func (r *rect) area() int {
	r.width += 2

	return r.width * r.height
}

// Методы могут принимать как указатели, так и значения. Вот пример со значением.
func (r rect) perim() int {
	r.width += 2

	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 5, height: 10}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())
	fmt.Println("width:", r.width)

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
	fmt.Println("width:", rp.width)
}
