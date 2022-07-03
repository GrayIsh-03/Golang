package main

import (
	"fmt"
	"golang/pkg/calculation"
	"golang/pkg/task"
	"log"
)

func main() {

	/*message := quote.Hello()
	fmt.Println(message)*/

	calculation.RocketSpeed()

	date := task.Date{}
	err := date.SetYear(2001)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
}
