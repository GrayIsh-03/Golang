// AverageArg в отличии от Average получает аргументы переданные текущей программе при запуске
// в пакетную переменную os.Args, которая содержит сегмент строк, не из файла, а из командной строки.
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Получение сегмента строк со всеми элементами os.Args, кроме первого(индекс [0]).
	arguments := os.Args[1:]
	var sum float64 // объявляем переменную для хранения суммы
	// Обрабатываем каждый аргумент командной строки.
	for _, argument := range arguments {
		// Строка преобразуется в float64.
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	// Длина сегмента arguments используется как количество значений данных.
	sampleCount := float64(len(arguments))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount) // Вычисление среднего значения и его вывод.
}
