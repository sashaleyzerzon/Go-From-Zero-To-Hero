package main

import "fmt"

func main05() {
	var name string
	var age int

	fmt.Println("Enter your name:")
	fmt.Scan(&name)
	fmt.Println("Enter age:")
	fmt.Scan(&age)
	fmt.Printf("Your name is %s and your age is %d", name, age)
}
