package main

import "fmt"

func main() {
	var name  string = "Nabheet"
	//short hand
	name2 := "Nabheet 2"
	size := 1.30
	var age int = 37
	fmt.Println(name,name2, age)
	fmt.Printf("%T\n",size)
	fmt.Printf("%T\n",age)
}