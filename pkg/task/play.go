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

// GuessNum this is game where player guessed random number between 1 and 100 to 10 guees left
func GuessNum() {
	seconds := time.Now().Unix() //получаем текущую дату и время в формате целого числа
	rand.Seed(seconds)           //функция генератора случайных чисел
	target := rand.Intn(100) + 1 //теперь генерируемые значения будут разные при каждом запуске
	fmt.Println("I've chosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")
	reader := bufio.NewReader(os.Stdin) //создаём буфферизированное средство для ввода с клавиатуры
	var succes bool

	for guesses := 0; guesses <= 10; guesses++ {
		fmt.Println("you have", 10-guesses, "guesses left.")
		fmt.Println("make a guesses:")
		input, err := reader.ReadString('\n') //прочитать данные, введённые пользователем до нажатия Enter
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)  // Удаление символа новой строки и почих служебных символов
		guess, err := strconv.Atoi(input) //входная строка преобразуется в целое число
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
