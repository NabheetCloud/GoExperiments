package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getsum(num1 int, num2 int) int {
	return num1 + num2
}
func main() {
	fmt.Println(greeting("Nabheet"))
	fmt.Println(getsum(1, 2))
}
