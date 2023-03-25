package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main06() {
	fmt.Println("The time now is", time.Now())
	fmt.Println("Random number between 0-100:", rand.Intn(100))
	fmt.Println("Square root of 9 is:", math.Sqrt(9.0))
}
