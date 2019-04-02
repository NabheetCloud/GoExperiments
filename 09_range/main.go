package main

import "fmt"

func main() {
	ids := []int{1, 2, 321, 21, 11}
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	for id := range ids {
		fmt.Printf(" ID: %d\n", id)
	}
	for _, id := range ids {
		fmt.Printf(" ID: %d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	email := map[string]string{
		"nabh": "test@test,com", "nabh1": "test@com"}
	for k, v := range email {
		fmt.Printf("%s: %s\n", k, v)
	}

}
