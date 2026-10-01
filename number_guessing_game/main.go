package main

import (
	"fmt"
	"math/rand"
)

func welcomeMessage() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	// maybe in future dynamic chance count
	fmt.Println("You have multiple chances to guess the correct number.")
	fmt.Println()
}

func selectDifficulty() (level, chances int) {
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")
	fmt.Println()

	var selection int
	fmt.Print("Enter your choise: ")
	fmt.Scan(&selection)
	fmt.Println()

	return selection, returnDifficulty(selection)

}

func returnDifficulty(level int) (chances int) {
	difficultyNumber := map[int]int{
		1: 10,
		2: 5,
		3: 3,
	}
	return difficultyNumber[level]
}

func returnDifficultyString(level int) string {
	difficultyNumber := map[int]string{
		1: "Easy",
		2: "Medium",
		3: "Hard",
	}
	return difficultyNumber[level]
}

func gameLoop(level, chances int) {

	// apparently this works
	randomNumber := rand.Intn(100) + 1

	levelString := returnDifficultyString(level)
	fmt.Printf("Great! You have selected the %s difficulty level.\n", levelString)
	fmt.Println("Let's start the game!")

	guessCorrectFlag := false

	for i := range chances {
		var guess int
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)

		if randomNumber == guess {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", i+1)
			guessCorrectFlag = true
			break

		} else if randomNumber < guess {
			fmt.Printf("Incorrect! The number is less than %d\n", guess)

		} else {
			fmt.Printf("Incorrect! The number is greater than %d\n", guess)
		}
	}

	if !guessCorrectFlag {
		fmt.Println("You ran out of chances")
	}

}

func main() {

	welcomeMessage()
	level, chances := selectDifficulty()
	gameLoop(level, chances)

}
