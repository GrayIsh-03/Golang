package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	// запускаем процесс вызова scanDirectory для каталолга верхнего уровня
	err := ScanDirectory("C:/Users/hp/Documents/GitHub/Golang")
	if err != nil {
		log.Fatal(err)
	}
	count(1, 3)
}

func count(start int, end int) {
	fmt.Println(start)
	if start < end {
		count(start+1, end)
	}
	fmt.Printf("Returning from count (%d, %d) call\n", start, end)
}

func ScanDirectory(path string) error {
	fmt.Println(path) // выводит текущую директорию
	// Получает сегмент с элементами, представляющими содержимое «path».
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		// Соединяет путь каталога и имя файла через символ \.
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() { // проверка на подкаталог
			// рекурсивно вызывает ScanDirectory, но на этот раз с путём подкаталога
			err := ScanDirectory(filePath)
			if err != nil {
				return err
			}
		} else {
			// если это файл, то вывести его имя с путём
			fmt.Println(filePath)
		}
	}
	return nil

}
