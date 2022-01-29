package main

import "fmt"

func main () {
	nums := []int{2, 3, 4, 5}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}

	// range для карт перебирает пары ключ/значение.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range может перебирать только ключи в карте
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range для строк перебирает кодовые точки Unicode.
	// Первое значение - это начальный байтовый индекс руны, а второе - сама руна.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
