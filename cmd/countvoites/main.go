package main

import (
	"fmt"
	"golang/pkg/fileprcss"
	"log"
)

func main() {
	lines, err := fileprcss.DataFileStr("C:/Users/hp/Documents/GitHub/voites.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
