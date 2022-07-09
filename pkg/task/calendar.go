package task

import (
	"fmt"
	"golang/internal/deftype"
	"log"
)

func CalenDate() {

	//* Опосредственное обращение
	// К неэкспортируемым переменным, полям структур, функциям, методам и т. д.
	// можно обращаться из экспортируемых функций и методов того же пакета.
	date := deftype.Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)

	/* date содержит значение типа Date. Обратиться к полям нельзя, т к они не экспортируюся
		при обращении к ним, как показано ниже будет ошибка.
	date.year = 2019 // защита от прямого присваивания
	date = deftype.Date{year: 2019, month: 3, day: 31} // защита от инициализации полей в литерале структуры
	fmt.Println(date)
	*/

	fmt.Println(date.Year()) // get-method for output unexported field
}
