package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	var guess int

	fmt.Println("Guess the number between 1 and 100!")

	for attempts := 1; ; attempts++ {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)

		if guess > target {
			fmt.Println("Too high!")
		} else if guess < target {
			fmt.Println("Too low!")
		} else {
			fmt.Printf("You guessed it in %d attempts!\n", attempts)
			break
		}
	}
}
