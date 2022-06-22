package main

import (
	"fmt"
	"golang/pkg/calculation"
)

func main() {

	/*message := quote.Hello()
	fmt.Println(message)*/

	calculation.RocketSpeed()

	fmt.Println(calculation.Sum(9, 1, 2, 4))
	fmt.Println(calculation.Sum(7))

}
