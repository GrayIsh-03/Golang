package minitask

import (
	"fmt"
	"math/rand"
	"time"
)

//func расчитывает сколько длится полет на Марс и расстояние до Марса в случайном диапозоне
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

// func расчитывает скорость ракеты необходимую для преодоления дистанции за определённое количество дней
func RocketSpeed() {
	const hoursPerDay = 24
	distance := 56000000
	dayOnRoad := 28

	rcktSpd := distance / dayOnRoad / hoursPerDay
	fmt.Printf("rocket speed is %v km/h\n", rcktSpd)
}

// выбор локации через условие
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

// func реализует шанс 1 к 10, что запуск будет отменён
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

}
