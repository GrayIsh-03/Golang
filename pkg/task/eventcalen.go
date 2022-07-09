// Для календаря назначается событие с определённой датой
package task

import (
	"fmt"
	"golang/internal/deftype"
	"log"
)

func EventCalen() {

	event := deftype.Event{}
	err := event.SetTitle("Mom's birthday")
	if err != nil {
		log.Fatal(err)
	}
	// exported methods of defined type Date
	err = event.SetYear(1965)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(7)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(9)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
}
