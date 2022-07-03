package pkg

import "fmt"

// объявление определяемго типа происходит на уровне пакета, в качестве
// базового типа используется тип struct
type tip struct {
	field1 string
	field2 bool
	field3 float64
}

func Base() {
	// краткое объявление переменной, где литерал структуры создаёт и одновременно инициализирует поля
	myStructure := tip{field1: "forth", field2: true, field3: 456}
	fmt.Println(myStructure)
}
