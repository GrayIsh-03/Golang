package main

import (
	"Golang/keyboard/readfile"
	"Golang/task"
	minitask "Golang/task/mini"
	"fmt"
)

func main() {

	/*message := quote.Hello()
	fmt.Println(message)*/

	task.RoundErr()
	minitask.Average()
	readfile.ReadFile()
	fmt.Println("=============================")

}
