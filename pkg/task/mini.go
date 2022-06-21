package task

import (
	"fmt"
	"math/rand"
	"time"
)

// func ThreeLocation реализует выбор локации через условие
func ThreeLocation() {
	var location = "not defined"

	if location == "cave" {
		fmt.Println("There is a very dark cave, where bats live")

	} else if location == "mountain" {
		fmt.Println("You climbed the highest peak of Mount Everest")

	} else if location == "entrance" {
		fmt.Println("This entrance leads to the dungeon")

	} else {
		fmt.Println("Location dont defined")
	}
}

// func LaunchFall реализует шанс 1 к 10, что запуск будет отменён
func LaunchFall() {
	var count int16
	for count = 10; count >= 0; count-- {
		fall := rand.Intn(10)
		if fall == 5 {
			fmt.Println("Launch of rocket canceled")
			break

		} else if count == 0 {
			fmt.Println("Launch is succesfull")

		} else {
			fmt.Println(count)
			time.Sleep(time.Second)
		}
	}
	/*
	   var count = 10
	   	for count > 0 {
	   		fmt.Println(count)
	   		fallChance := rand.Intn(21)

	   		if fallChance == 2 { // chance 1 in 20
	   			break
	   		}

	   		time.Sleep(time.Second)
	   		count--
	   	}

	   	switch count {
	   	case 0:
	   		fmt.Println("Launch!")
	   	default:
	   		fmt.Println("Launch canceled!")
	   	}
	*/
}

// покупка билетов до Марса, формируется таблица, чем больше скорость тем больше цена
// при путешествии туда и обратно цена увеличивается вдвое.
func TicketsToMars() {
	const (
		Distance = 62100000
		Hours    = 24
	)

	fmt.Printf("Spaceline %12v %13v %8v\n===============================================\n", "Days", "Trip type", "Price")

	for count := 0; count < 10; count++ {
		// spdShip range for 16 to 30 km/h
		spdShip := 16000 + rand.Intn(14001)
		// price range for 36 to 50 millions $
		price := spdShip/1000 + 20
		day := Distance / spdShip / Hours
		spaceStaion := "Space Adventures"
		tripType := "One-way"

		switch st := rand.Intn(3); st {
		case 0:
			spaceStaion = "SpaceX          "
		case 1:
			spaceStaion = "Virgin Galactic "
		}

		switch tt := rand.Intn(2); tt {
		case 0:
			tripType = "Round-trip"
			price *= 2
		}

		fmt.Printf("%v %5v %13v    $%5v\n", spaceStaion, day, tripType, price)
	}
}

// В данном примере мы выполняем цикл через строку, которая содержит
// различные символы. Выводятся все, кроме пробелов и символов табуляции и новой строки.
func EnumerationCharacter() {
	w := "a b c\td\nefg hi"

	for _, e := range w {

	l:
		switch e {
		case ' ', '\t', '\n':
			break l
		default:
			fmt.Printf("%c\n", e) // %с outputs a character according to the Unicode code
		}
	}
}
