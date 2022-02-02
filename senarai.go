package main

import "fmt"

func main() {
	var a [3]int
	a[0] = 1
	fmt.Println(a[0])

	fmt.Println(len(a))

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	var person = []map[string]interface{}{
		{"name": "Wick", "age": 23},
		{"name": "Jono", "age": 31},
		{"name": "Sino", "age": 43},
	}
	fmt.Println(person[0]["name"])
	fmt.Println(len(person))
}
