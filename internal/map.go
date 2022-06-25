package internal

import "fmt"

func KeyValueMap() {
	data := []string{"a", "c", "e", "a", "e"}
	// создание карты в переменной counts, ключи содержат нулевые значения
	// сама переменная значение nil
	counts := make(map[string]int)
	for _, item := range data {
		/* карта заполняется ключами item и значениями, которые изначально нулевые,
		но при использовании инкремента увеличивают свое значение на единицу
		каждый новый ключ (а с е) получает значение 1, если ключ повторяется,
		то происходит увелечение его предыдущего значения на еденицу, т.о.
		можно определить количество повторяющихся элементов в сегменте data*/
		counts[item]++
	}
	letters := []string{"a", "b", "c", "d", "e"}
	for _, letter := range letters {
		/* переменной count присваивается значение из карты counts по ключу
		letter, если letter == item, то записывается значение соответствующее
		ключу item из цикла выше, если нет, то нулевое значение для count
		и логическое значение false для ок, говорящее о том, что ключу не было
		присвоено значение*/
		count, ok := counts[letter]
		if !ok {
			fmt.Printf("%s: not found\n", letter)
		} else {
			fmt.Printf("%s : %d\n", letter, count)
		}
	}
}
