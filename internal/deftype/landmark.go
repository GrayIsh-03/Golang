// Определяет тип LandMark на базе структуры со встроенным типом Coordinates
package deftype

import "errors"

type LandMark struct {
	name        string // Поле name не должно экспортироваться, чтобы оно было инкапсулировано.
	Coordinates        // Встраивание с помощью анонимного поля из deftype/geo.go
}

// get-method for field "name"
func (l *LandMark) Name() string {
	return l.name
}

// set-method for field "name"
func (l *LandMark) SetName(name string) error {
	if name == "" {
		return errors.New("invalid name")
	}
	l.name = name
	return nil
}
