package main

import "fmt"

type person struct {
	name string
}

func main() {

	p := person{"Richard"}
	fmt.Println(p)
	rename(&p)
	fmt.Println(p)
	renamer(p)
	fmt.Println(p)

	var nama string = "agusta"
	var point *string = &nama
	fmt.Println(nama)
	fmt.Println(point)
	*point = "sama"
	fmt.Println(*point)
	fmt.Println(nama)
}

func rename(p *person) {
	p.name = "test"
}

func renamer(p person) person {
	p.name = "asal muni"
	return p
}
