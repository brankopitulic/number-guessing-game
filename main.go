package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num int
	fmt.Println("Guess number game between 0 and 10!")
	fmt.Println("Input number:")

	randomNum := rand.Intn(10)

	fmt.Scan(&num)

OuterLoop:
	for i := 0; i < 3; i++ {
		switch {
		case num == randomNum && i <= 2:
			fmt.Println("Congrats")
			fmt.Println(randomNum)
			break OuterLoop
		case num < randomNum && i < 2:
			fmt.Println("It's to low")
			fmt.Println(randomNum)
			newTry(&num)
			continue OuterLoop
		case num > randomNum && i < 2:
			fmt.Println("It's to high")
			fmt.Println(randomNum)
			newTry(&num)
			continue OuterLoop
		case i == 2:
			fmt.Printf("You lost, random number was: %d", randomNum)
		}
	}

}

func newTry(num *int) {
	var newNum int
	fmt.Println("Try again:")

	fmt.Scan(&newNum)

	*num = newNum
}
