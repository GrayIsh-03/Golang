package calculation

import (
	"fmt"
	"golang/pkg/fileprcss"
	"log"
)

//func Average вызывает func DataFile, от которой получает массив значений из файла
//через цикл for range производит их суммирование и высчитывает среднее значение
func Average() {
	//загружает файл по указанному пути, разбирает содержащиеся в нём числа
	// и сохраняет в массив
	numbers, err := fileprcss.DataFileFlo("C:/Users/hp/Documents/GitHub/data.txt")
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
