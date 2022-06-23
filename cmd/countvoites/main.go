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
	var names []string // сегмент для хранения имён кандидатов
	var counts []int   // сегмент с количеством вхождения каждого имени

	for _, line := range lines {
		matched := false

		for i, name := range names {

			if name == line { // если строка совпадает с текущим именем
				counts[i]++    // увеличиваем соответствующий счётчик
				matched = true // устанавливает признак обнаруженного совпадения
			}

		}

		if matched == !true { // если совпадение не найдено
			names = append(names, line) // добавить его как новое имя в сегмент
			counts = append(counts, 1)  // и добавить новый счётчик (текущая строка станет 1-ым вхождением)
		}

	}

	for i, name := range names {
		// Выыести каждый элемент из сегмента names и соответствующий элемент из сегмента counts
		fmt.Printf("%s: %d\n", name, counts[i])
	}

}
