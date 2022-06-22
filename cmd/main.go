package main

import (
	"fmt"
	"golang/pkg/calculation"
)

func main() {

	/*message := quote.Hello()
	fmt.Println(message)*/

	calculation.RocketSpeed()

	fmt.Println(calculation.InRange(1, 20, 3, 78, 34, 3, 15, 56))

}
