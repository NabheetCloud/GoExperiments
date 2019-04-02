package main

import "fmt"

func main() {
	emails := make(map[string]string)
	emails["nabheet"] = "test@test.com"
	emails["nabheet1"] = "test@test.com1"
	emails["nabheet2"] = "test@test.com2"

	fmt.Println(emails)
	fmt.Println(emails["nabheet"])
	fmt.Println(len(emails))

	delete(emails, "nabheet")
	fmt.Println(emails)
	email := map[string]string{
		"nabh": "test@test,com", "nabh1": "test@com"}
	fmt.Println(email)
}
