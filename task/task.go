package task

import (
	"Golang/keyboard"
	"Golang/paintneed"
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

// func определяет сколько нужно краски на каждую стену и сколько в общем
func PaintSum() {
	var total float64
	amount, err := paintneed.PaintNeeded(12.2, 3)
	if err != nil {
		fmt.Println(err)
	}
	total += amount
	amount, err = paintneed.PaintNeeded(8.3, -2.5)
	if err != nil {
		fmt.Println(err)
	}
	total += amount
	amount, err = paintneed.PaintNeeded(12, 8)
	if err != nil {
		fmt.Println(err)
	}
	total += amount
	fmt.Printf("Total: need %0.2f a liters paint\n", total)
}

func Pass() {
	fmt.Print("Enter your a grade:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(input)
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)

	if err != nil {
		log.Fatal(err)
	} else if grade >= 60 {
		fmt.Println("Your passed")
	} else {
		fmt.Println("Your failed")
	}

	fileInfo, err := os.Stat("my.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
	fmt.Print(fileInfo.ModTime(), fileInfo.Sys())

}

func LeapYear() {

	for count := 10; count > 0; count-- {

		year := 2018 + rand.Intn(20)
		mounth := rand.Intn(12) + 1
		leapYear := year%400 == 0 || (year%4 == 0 && year%100 != 0)
		dayInMounth := 31

		if leapYear {
			fmt.Printf("%10v will leap year\n", year)
		} else {
			fmt.Printf("%10v will non-leap year\n", year)
		}

		switch mounth {
		case 2:
			dayInMounth = 28
			if leapYear {
				dayInMounth = 29
			}
		case 4, 6, 9, 11:
			dayInMounth = 30
		}

		day := rand.Intn(dayInMounth)
		fmt.Println(year, mounth, day)

	}
}

//функция сообщает сдал ли пользователь экзамен
func PassFail() {
	fmt.Println("Enter a grade:")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade < 60 {
		status = "failling"
	} else {
		status = "passing"
	}
	fmt.Println("A grade of", grade, status)
}

func ToCelsius() {
	fmt.Print(" Enter a temperature in Fahrenheit:")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", celsius)
}
