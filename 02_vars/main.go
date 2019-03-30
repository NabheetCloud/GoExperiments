package main

import "fmt"

func main() {
	// different ways for variable declarations
	var name string = "Nabheet"
	//short hand
	name2 := "Nabheet 2"
	size := 1.30
	lastname, lastname2 := "Madan", "Madan2"
	var age int = 37
	fmt.Println(name, name2, age)
	fmt.Println(lastname, lastname2)
	fmt.Printf("%T\n", size)
	fmt.Printf("%T\n", age)
}
