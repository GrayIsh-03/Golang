// package main рассматривает два варианта подсчёта повторяющихся значений
// с помощью сегментов и карты
package main

import (
	"fmt"
	"golang/pkg/fileprcss"
	"log"
	"sort"
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
	countVoites()
}

// func countVoites реализует подсчёт повторяющихся значений строкового типа
// с помощью объявления и создания карты
func countVoites() {
	lines, err := fileprcss.DataFileStr("C:/Users/hp/Documents/GitHub/voites.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int) // объявляет карту: ключи - имена кандидатов, значения счётчик голосов
	for _, line := range lines {
		counts[line]++ // увеличивает счётчик голосов для текущего кандидата
	}
	var names []string         // создаём пустой сегмент для хранения ключей из карты counts
	for name := range counts { // записываем ключи из counts в name, в результате они
		// станут значениями для сегмента names, пустой идентификатор не используется
		// заполняем сегмент names
		names = append(names, name)
	}
	sort.Strings(names) // сортируем значения стркового типа сегмента names в алфавитном порядке
	for _, name := range names {
		// сегмент names содержит значения, которые соответствуют ключам в counts
		// name получает значения/ключи в алфавитном порядке
		fmt.Printf("Votes a %s of %d\n", name, counts[name])
		// fmt выводит name как значение сегмента, а [name] как значение по ключу из карты
	}
}
