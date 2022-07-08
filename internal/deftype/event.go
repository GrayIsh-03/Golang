// рассмотрен пример повышения методов
package deftype

import (
	"errors"
	"unicode/utf8"
)

// insert type Date to place deftype/calendar.go
type Event struct {
	title string // объявляем unexported
	Date         // повышение полей Date до уровня Event не произойдёт, т.к. поля unexported
	// но set and get methods Date exported, они повышаются до уровня Event
	// и к полям можно обращаться через них
}

// get-method Event.Title
func (e *Event) Title() string {
	return e.title
}

// set-method Event.Title
func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 { // подсчёт количества рун в строке
		return errors.New("invalid title") // если рун > 30, возвращается ошибка
	}
	e.title = title
	return nil
}

/*
Теперь мы можем вызывать методы Title и SetTitle прямо для Event, а методы
для присваивания year, month и day так, как если бы они принадлежали Event.
На самом деле они определяются для Date, но нам не обязательно об этом помнить.
*/
