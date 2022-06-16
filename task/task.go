package task

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

// this a function visually shows rounding errors floating point numbers
func RoundErr() {
	var total float64
	for i := 1; i <= 10; i++ {
		total += 0.10
	}
	fmt.Println(total) // will display the result 0.9999999999999999

	numFlo := 0.1
	numFlo += 0.2
	fmt.Println(numFlo) // will display the result 0.30000000000000004
	// для сравнения чисел с плавающей запятой убедитесь, что абсолютная
	// разница между двумя числами не слишком велика
	fmt.Println(numFlo == 0.3)                 // false
	fmt.Println(math.Abs(numFlo-0.3) > 0.0001) // false
}

// Эта func выбирает случайные значения и суммирует их с выводом текущего значения
// суммирование продолжается до заданного значения
func MoneyRand() {
	var box float64
	for {
		i := rand.Intn(3)
		switch i {
		case 0:
			box += 0.05
		case 1:
			box += 0.10
		case 2:
			box += 0.25
		}
		fmt.Printf("In your moneybox %05.2f $\n", box)
		if box >= 20 {
			break
		}
	}
}
