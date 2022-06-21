package calculation

import (
	"fmt"
	"golang/pkg/fileprcss"
	"log"
	"os"
	"strconv"
)

//func Average вызывает func DataFile, от которой получает массив значений из файла
//через цикл for range производит их суммирование и высчитывает среднее значение
func Average() {
	//загружает файл по указанному пути, разбирает содержащиеся в нём числа
	// и сохраняет в массив
	numbers, err := fileprcss.DataFile("C:/Users/hp/Documents/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64
	//с помощью <_> искл index, за каждый цикл в number записывается значение из массива numbers
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers)) //присваиваем количество элементов массива
	average := sum / sampleCount         //высчитывается среденее значение всех элементов массива
	fmt.Printf("Average: %0.2f\n", average)
}

// func AverageArg в отличии от Average получает аргументы переданные текущей программе при запуске
// в пакетную переменную os.Args, которая содержит сегмент строк, не из файла, а из командной строки.
func AverageArg() {
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
