package task

import (
	"fmt"
	"golang/internal/deftype"
)

func defaultSubscriber(name string) *deftype.Subscraiber {
	var s deftype.Subscraiber
	s.Name = name
	s.Rate = 5.99
	s.Active = true
	return &s
}

func applyDiscount(s *deftype.Subscraiber) {
	s.Rate = 4.99
}

func printInfo(s *deftype.Subscraiber) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Monthly rate:", s.Rate)
	fmt.Println("Active?", s.Active)
}

func Subscriber() {

	subscraiber1 := defaultSubscriber("Aman Singh")
	applyDiscount(subscraiber1)
	printInfo(subscraiber1)
	subscraiber2 := defaultSubscriber("Beth Rayn")
	printInfo(subscraiber2)

	var employee deftype.Employee
	employee.Name = "Dmitry Axe"
	employee.Salary = 234.2
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)

	/* //* Создание вложенной структуры. Вариант первый: создаём отдельную структуру
	   adsress, а затем используем её для заполнения всего поля HomeAddress структуры
	   subscriber

	address := deftype.Address{Street: "Vakarina", City: "Ulan-Ude",
		State: "Buraytiya", PostalCode: "3012600",
	}
	subscraiber := deftype.Subscraiber{Name: "Dmitriy Sword"}
	subscraiber.HomeAddress = address
	fmt.Println(subscraiber.HomeAddress)
	*/

	/* //* Создание вложенной структуры. Вариант второй: присвоение полям внутренней
	   структуры значений через внешнюю структуру. Внешняя стуктура subscriber уже
	   содержит поле HomeAddress, которое имеет тип вложенной структуры Adreess и
	   заполнено нулевыми значениями. Если subscriber — переменная, содержащая структуру
		Subscriber, то конструкция subscriber.HomeAddress дает вам структуру Address, хотя
		вы и не задали HomeAddress явно. Это позволяет использовать «сцепленные» операторы
		«точка» для обращения к полям структуры Address. subscriber.HomeAddress.City = ""
	*/
	subscriber := deftype.Subscraiber{Name: "Dmitry Knife"}
	// subscriber.HomeAddress.Street = "Vakarina 42" вариант обращения к неанонимному полю
	subscriber.City = "Ulan-Ude" // обращение к анонимному полю
	subscriber.State = "Saha Yakutiya"
	subscriber.PostalCode = "680780"
	fmt.Printf("%#v\n\n", subscriber)

	employee.Street = "Renfoks 12"
	employee.City = "Oyimyakon"
	employee.State = "Saha Yakutiya"
	employee.PostalCode = "680790"
	fmt.Printf("%#v\n", employee)
	fmt.Println(employee.State)
}
