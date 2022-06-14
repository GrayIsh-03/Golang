package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	message := quote.Hello()
	fmt.Println(message)
}
