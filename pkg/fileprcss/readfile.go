package fileprcss

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFile() {
	file, err := os.Open("C:/Users/hp/Documents/data.txt") // файл данных открывается для чтения
	if err != nil {
		log.Fatal(err) // если при открытии файла произошла ошибка, сообщить о ней и завершить работу
	}
	scanner := bufio.NewScanner(file) // for file create new value Scanner
	// цикл выполняется до того, как будет достигнут конец файла, а scanner.Scan return false
	for scanner.Scan() { //read string of file
		fmt.Println(scanner.Text()) // out string, если строки не переносятся, см. func Println
	}
	err = file.Close() // close file for освобождения ресурсов
	if err != nil {
		log.Fatal(err) // если при закрытии файла произошла ошибка, сообщить о ней и завершить работу
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err()) // если при сканировании файла произошла ошибка, сообщить о ней и завершить работу
	}
}
