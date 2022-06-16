package main

import (
	"Golang/paintneed"
	"Golang/task"

	"fmt"
)

func main() {

	/*message := quote.Hello()
	fmt.Println(message)*/

	// func определяет сколько нужно краски на каждую стену и сколько в общем
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

	task.RoundErr()
	task.MoneyRand()
}
