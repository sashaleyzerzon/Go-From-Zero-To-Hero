package main

import "fmt"

func main02() {
	const myName string = "Sasha"
	var myAge int = 25

	// Exercise #1
	fmt.Println("My name is", myName, "and my age is", myAge)

	// Exercise #2 - answer
	fmt.Println("It will not work because there is a declared variable, which is unused.")

	// Exercise #3 - answer
	fmt.Println("Squares:2, Circles:0 \n Squares:1, Circles:7")
	fmt.Println("EDIT: it will fail and I didn't notice that Squares is a const!")
}
