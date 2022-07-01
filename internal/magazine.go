// в этом пакете рассмотрен пример экспортирования определяемого типа на
// базовом типе структуры и экспорт полей структуры.
package internal

type Subscraiber struct {
	Name        string
	Rate        float64
	Active      bool
	HomeAddress Address
}

type Employee struct {
	Name        string
	Salary      float64
	HomeAddress Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
