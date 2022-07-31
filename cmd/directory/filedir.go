package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	// запускаем процесс вызова scanDirectory для каталолга верхнего уровня
	/*//* code for func ScanDirectory
	err := ScanDirectory("C:/Users/hp/Documents/GitHub/Golang")
	if err != nil {
		log.Fatal(err)
	}*/

	//* code for func ScanDirectoryPanic
	ScanDirectoryPanic("C:/Users/hp/Documents/GitHub/Golang")
	count(1, 3)
}

// recursive function call example
func count(start int, end int) {
	fmt.Println(start)
	if start < end {
		count(start+1, end)
	}
	fmt.Printf("Returning from count (%d, %d) call\n", start, end)
}

// Выводит содержимое каталога с помощью рекурсивной функции
func ScanDirectory(path string) error {
	fmt.Println(path) // выводит текущую директорию
	// Получает сегмент с элементами, представляющими содержимое «path».
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	// Функция перебирает сегмент значений FileInfo, возвращенных ReadDir
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

func ScanDirectoryPanic(path string) { // значение ошибки уже не нужно
	fmt.Println(path) // выводит текущую директорию
	// Получает сегмент с элементами, представляющими содержимое «path».
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	// Функция перебирает сегмент значений FileInfo, возвращенных ReadDir
	for _, file := range files {
		// Соединяет путь каталога и имя файла через символ \.
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() { // проверка на подкаталог
			// рекурсивно вызывает ScanDirectory, но на этот раз с путём подкаталога
			ScanDirectory(filePath)
			// cохранять или проверять воз-вращаемое зна-чение ошибки уже не нужно
		} else {
			// если это файл, то вывести его имя с путём
			fmt.Println(filePath)
		}
	}

}
