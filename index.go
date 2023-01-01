
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var userOption = readUserOption()
	printWinner(userOption)
}

func readUserOption() string {
    fmt.Println("Insert 'p' to paper, 'r' to rock or 's' to scissors")
	var selection string
	for (selection != "p" && selection != "r" && selection != "s") {
		fmt.Scanln(&selection)
	}
	return selection
}

func printWinner(userChoice string) {
	var computerChoice string
	computerChoice = getComputerChoice()
	
	if (((computerChoice == "p") && (userChoice == "s")) || ((computerChoice == "r") && (userChoice == "p")) || ((computerChoice == "s") && (userChoice == "r"))) {
		fmt.Println("User win")
	} else if (computerChoice == userChoice) {
		fmt.Println("Draw")
	} else {
		fmt.Println("Computer win")
	}

	fmt.Print("User selected: ")
	fmt.Println(userChoice)
	fmt.Print("Computer selected: ")
	fmt.Println(computerChoice)
}

func getComputerChoice() string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	computerIntChoice := random.Intn(3)
	if (computerIntChoice == 0) {
		return "p"
	} else if (computerIntChoice == 1) {
		return "r"
	}
	return "s"
}