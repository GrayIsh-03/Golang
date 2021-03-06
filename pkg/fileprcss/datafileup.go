// Рассмотрен пример use defer
package fileprcss

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Открывает файл и возвращает указатель на него и значение обнаруженной ошибки
func OpenFile(filename string) (*os.File, error) {
	fmt.Println("Opening", filename)
	return os.Open(filename)
}

// закрывает файл
func CloseFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}

// func производит чтение из файла value float64 и записывает их в []float64
// use отложенный вызов defer для закрытия файла в случае возвращения ошибки
func DFileFlo(filename string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(filename)
	if err != nil {
		return nil, err
	}
	defer CloseFile(file) // используем отложенный вызов, func будет выполнена даже при наличии ошибок в коде
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}

// суммирует значения float64 из полученного сегмента
func SumFlo() {
	numbers, err := DFileFlo(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Sum: %0.2f\n", sum)
}
