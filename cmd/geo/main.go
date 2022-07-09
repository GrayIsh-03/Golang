package main

import (
	"fmt"
	"golang/internal/deftype"
	"log"
)

func main() {
	// location содержит значение типа LandMark, через него вызываем методы
	// для LandMark и встроенного Coordinates
	location := deftype.LandMark{}
	err := location.SetName("The Googleplex") // вызываем на уровне landMark
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLatitude(37.42) // вызываем на уровне landMark от встроенного типа Coordinates
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLongitude(-122.08)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(location.Name())      // Определяется от landMark
	fmt.Println(location.Latitude())  // Повышается от Coordinates
	fmt.Println(location.Longitude()) // Повышается от Coordinates
}
