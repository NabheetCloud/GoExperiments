package main

import "fmt"

func main() {
	var fruitArr [2]string
	fruitArr[0] = "a"
	fruitArr[1] = "b"

	fruitArr2 := [2]string{"ao", "bo"}
	fruitSlice := []string{"ao", "bo", "co"}
	fmt.Println(len(fruitArr))
	fmt.Println(fruitArr2)
	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[0:1])

}
