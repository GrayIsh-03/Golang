package main

import (
	"fmt"
	"golang/pkg/calculation"
)

type subscraiber struct {
	name   string
	rate   float64
	active bool
}

func main() {

	/*message := quote.Hello()
	fmt.Println(message)*/

	calculation.RocketSpeed()

	subscraiber1 := defaultSubscriber("Aman Singh")
	subscraiber1.rate = 4.99
	printInfo(subscraiber1)

	subscraiber2 := defaultSubscriber("Beth Rayn")
	printInfo(subscraiber2)
}

func printInfo(s subscraiber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}

func defaultSubscriber(name string) subscraiber {
	var s subscraiber
	s.name = name
	s.rate = 5.99
	s.active = true
	return s
}
