// в этом пакете рассмотрен пример экспортирования определяемого типа на
// базовом типе структуры и экспорт полей структуры.
package internal

type Subscraiber struct {
	Name   string
	Rate   float64
	Active bool
	// HomeAddress Address данное поле типа структуры содержит вложенную структуру Address
	Address // create an anonymous field
}

type Employee struct {
	Name   string
	Salary float64
	// HomeAddress Address
	Address // create an anonymous field для упрощения кода при обращении к поллям
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
