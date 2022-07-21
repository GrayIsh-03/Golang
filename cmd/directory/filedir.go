package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// Получает сегмент с элементами, представляющими содержимое «my_directory».
	files, err := ioutil.ReadDir("my_directory")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Directory:", file.Name())
		} else {
			fmt.Println("File:", file.Name())
		}
	}
}
