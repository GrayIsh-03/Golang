package fileprcss

import (
	"bufio"
	"os"
	"strconv"
)

// func DataFile открывает файл, считывает строки, преобразует их в float64
// и записывает их в массив значений и возвращает массив из трёх элементов или ошибку
func DataFile(filename string) ([]float64, error) {
	/*
		var numbers [3]float64 // объявление возвращаемого массива
		// открывает файл с переданным именем, передаём аргумент функции os.Open, а не путь в виде строки
		file, err := os.Open(filename)
		if err != nil {
			return numbers, err
		}
		i := 0 // переменная для хранения индекса, по которому должно выполняться присваивание
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			// строка прочитанная из файла преобразуется в float64
			// пустая строка выдаст ошибку parsing
			numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				return numbers, err
			}
			i++ //переход к следующему индексу массива
		}
		err = file.Close()
		if err != nil {
			return numbers, err
		}
		if scanner.Err() != nil {
			return numbers, scanner.Err()
		}
		// Если выполнение дошло до этой точки, значит,ошибок не было, поэтому программа
		//возвращает массив чисел и значение ошибки «nil».
		return numbers, nil
	*/
	// numbers по умолчанию содержит nil. append интерпретирует nil как пустой сегмент
	var numbers []float64 // реализуем сегмент для расширения массива новыми значениями
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// строка преобразуется в float64 и присваивается временной переменной
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number) // новое число присоединяется к сегменту
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, err
}
