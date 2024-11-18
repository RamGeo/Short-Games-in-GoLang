package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	choices := []string{"rock", "paper", "scissors"}
	rand.Seed(time.Now().UnixNano())

	var playerChoice string
	fmt.Print("Choose rock, paper, or scissors: ")
	fmt.Scan(&playerChoice)

	computerChoice := choices[rand.Intn(len(choices))]
	fmt.Println("Computer chose:", computerChoice)

	if playerChoice == computerChoice {
		fmt.Println("It's a draw!")
	} else if (playerChoice == "rock" && computerChoice == "scissors") ||
		(playerChoice == "paper" && computerChoice == "rock") ||
		(playerChoice == "scissors" && computerChoice == "paper") {
		fmt.Println("You win!")
	} else {
		fmt.Println("You lose!")
	}
}
