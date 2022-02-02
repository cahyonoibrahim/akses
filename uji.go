package main

import (
	"fmt"

	. "github.com/klauspost/cpuid/v2"
)

func main() {
	fmt.Println("PhysicalCores:", CPU.BrandName)

	fmt.Println("ThreadsPerCore:", CPU.ThreadsPerCore)
	//runtime.GOMAXPROCS
	var arSatu = [...]string{
		"satu",
		"dua",
		"tiga",
		"empat",
		"lima",
	}
	var arSatuBaru = arSatu[2:4]

	fmt.Println(arSatu)

	fmt.Println(arSatuBaru)

}
