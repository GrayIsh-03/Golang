package task

import (
	"errors"
)

type Date struct {
	// имена полей с нижнего регистра для запрета от экспортирования
	// т.о. достигается защита от ввода недействительных данных напрямую
	year  int
	month int
	day   int
}

// Set-method проверяет что бы значение year было больше еденицы, в противном случае ошибка
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

// Set-method проверяет что бы значение month лежало в пределах от 1 до 12, иначе ошибка
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

// Set-method проверяет что бы значение day лежало в пределах от 1 до 31, иначе ошибка
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}
