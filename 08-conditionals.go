package main

import (
	"fmt"
	"math/rand"
)

func main08() {
	nRandNum := rand.Intn(100-0) + 1

	if nRandNum == 50 {
		fmt.Println("The number is 50!")
	} else if nRandNum > 50 {
		fmt.Print(nRandNum, " is closer to 100")
		if nRandNum%2 == 0 {
			fmt.Print(" and it's even!\n")
		} else {
			fmt.Println()
		}
	} else if nRandNum < 50 {
		fmt.Println(nRandNum, "is closer to 0")
	}
}
