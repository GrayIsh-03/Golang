package main

import (
	"fmt"
	"golang/internal/geo"
	"golang/pkg/calculation"
	"log"
)

func main() {

	calculation.RocketSpeed()

	coordinates := geo.Coordinates{}

	err := coordinates.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}

	err = coordinates.SetLongitude(-1182.2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(coordinates.Latitude())
	fmt.Println(coordinates.Longitude())
}
