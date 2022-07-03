// package calculation содержит функции, которые на основе заданных аргументов
// или переменных производят математические вычисления и отображают или
// возвращают результат
package calculation

import (
	"fmt"
	"golang/pkg/keyboard"
	"log"
	"math"
	"math/rand"
	"time"
)

type Liters float64
type Gallons float64
type Milliliters float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliletrs() Milliliters {
	return Milliliters(g * 3785.41)
}

//func VarpEngine расчитывает сколько длится полет на Марс и расстояние до Марса в случайном диапозоне
func VarpEngine() {
	const hoursPerDay = 24
	var (
		speed    = 100800
		distance = 96300000
	)

	day := distance / speed / hoursPerDay
	fmt.Printf("Ours astronauts will fly to the Mars in %15v days \n", day)
	// random distance from 56 000 000 to 401 000 000 km
	randDistance := 56000000 + rand.Intn(345000001)
	fmt.Println(randDistance, "km this random distance from the Earth to the Mars")
}

// func RocketSpeed расчитывает скорость ракеты необходимую для преодоления дистанции за определённое количество дней
func RocketSpeed() {
	const hoursPerDay = 24
	distance := 56000000
	dayOnRoad := 28

	rcktSpd := distance / dayOnRoad / hoursPerDay
	fmt.Printf("rocket speed is %v km/h\n", rcktSpd)
}

// func ToCelsius import GetFloat <- pkg/keyboard/keyboar.go
// func ToCelsius получает пользовательский ввод с клавиатуры, преобразованный в число
// и вычисляет значение градусов из Фаренгейта в Цельсия
func ToCelsius() {
	fmt.Print(" Enter a temperature in Fahrenheit:")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", celsius)
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

// func MoneyBox реализует цикл в котором, в переменную box суммируются случайным
// образом каждый раз одно из 3-х значений до достижения заданного условия
// после чего результат форматируется к определённому виду и отображается,
// используется задержка по времени
func MoneyBox() {
	var box int32
	for box < 2000 {

		switch i := rand.Intn(3); i {
		case 0:
			box += 5
		case 1:
			box += 10
		case 2:
			box += 25
		}

		dollars := box / 100
		cents := box % 100
		fmt.Printf("In your money box $%d.%02d\n", dollars, cents)
		time.Sleep(time.Millisecond)
	}
}

// func Maximum получает любое количество аргументов float64 и возвращает
// наибольшее значение среди всех переданных.
func Maximum(numbers ...float64) float64 {
	max := math.Inf(-1) // присваивается очень низкое значение
	for _, number := range numbers {
		if number > max {
			max = number // находит наибольшее значение среди аргументов
		}
	}
	return max
}

// func InRange возвращает сегмент значений лежащих в указанном диапозоне
// от min до max
func InRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		// если аргумент не ниже min и не выше max
		if number >= min && number <= max {
			// то он добавляется в сегмент result
			result = append(result, number)
		}
	}
	return result
}

// func суммирует передаваемые аргументы и возвращает их сумму
func Sum(numbers ...int) int {
	var sum int = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// func ConvGalLit конвертирует единицы измерения liters to gallons, milliliters to gallons, gallons to liters and milliliters
func ConvGalLit(g float64, l float64, ml float64) {
	soda := Liters(l)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())
	water := Milliliters(ml)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons())
	milk := Gallons(g)
	fmt.Printf("%0.3f gallons equals %0.3f liters\n%0.3f gallons equals %0.3f milliliters\n", milk, milk.ToLiters(), milk, milk.ToMilliletrs())

}
