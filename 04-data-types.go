package main

import (
	"fmt"
	"reflect"
)

func main04() {
	// Exercise 01
	fmt.Println("It will not work, var userName is uninitialized and it's data type is unstated")

	// Exersice 02
	fmt.Println("The content of the var userName, which in our case is \"user\" will be printed")

	// Exercise 03
	var userName string
	userName = "user"
	fmt.Println(userName)

	// Exercise 04
	var food = "Pizza"
	var slices = 4
	var pineappleOnPizza = true

	fmt.Println("Type of food var is", reflect.TypeOf(food))
	fmt.Println("Type of slices var is", reflect.TypeOf(slices))
	fmt.Println("Type of pineappleOnPizza var is", reflect.TypeOf(pineappleOnPizza))
	fmt.Printf("food is %T\nslices is %T\npineappleOnPizza is %T", food, slices, pineappleOnPizza)

}
