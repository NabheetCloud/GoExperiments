package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

func (p Person) greet() string {
	return "hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

func (p *Person) hasBirthday() {
	p.age++
}
func (p *Person) getMarried(spouseLastname string) {
	p.lastName = spouseLastname
}
func main() {
	person1 := Person{
		firstName: "Nabheet",
		lastName:  "madan",
		city:      "asa",
		gender:    "m",
		age:       10}

	fmt.Println(person1)
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)
	person1.hasBirthday()
	person1.getMarried("test")
	fmt.Println(person1.greet())
}
