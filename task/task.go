package task

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// GuessNum this is game where player guessed random number between 1 and 100 to 10 left
func GuessNum() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")
	reader := bufio.NewReader(os.Stdin)
	var succes bool

	for guesses := 0; guesses <= 10; guesses++ {
		fmt.Println("you have", 10-guesses, "guesses left.")
		fmt.Println("make a guesses:")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops, your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops, your guess was HIGHT.")
		} else {
			succes = true
			fmt.Println("Yes, your Guessed it!!!")
			break
		}
	}

	if !succes {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}

}
