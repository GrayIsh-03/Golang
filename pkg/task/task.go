package task

import (
	"bufio"
	"fmt"
	"golang/pkg/keyboard"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// func RoundErr visually shows rounding errors floating point numbers
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

// func LeapYear выводит 10 случайных дат в определённом формате
// при чём указывается количество дней для соответствующего месяца, а так же
// количество дней в феврале с учётом високосного года.
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

// func PassFail import GetFloat <- pkg/keyboard/keyboard.go
// func PassFail сообщает сдал ли пользователь экзамен, отличие от Pass в том,
// что часть кода вынесено в func GetFloat, который может использоваться в
// других пакетах (calculation -> ToCelsius), код становится более наглядным.
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
