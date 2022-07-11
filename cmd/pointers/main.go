package main

import "fmt"

type mySubscriber struct {
	name   string
	old    int
	rate   float64
	active bool
}

func main() {

	//! Pointers: all possible cases

	//* First case
	var myIntPointer *int // Declaring a variable containing a pointer to pointer type int
	var myInt int = 4
	myIntPointer = &myInt     // Pointer is assigned to a variable
	fmt.Println(myIntPointer) // prints the address of the memory location where the variable is stored

	//? читается как "значение по адресу myIntPointer"
	fmt.Println(*myIntPointer) // prints the value pointed to by the pointer
	*myIntPointer = 8          // The new value is assigned to the variable myInt
	fmt.Println(myInt)
	fmt.Println("=====================================================")
	//todo=====================================================================
	fmt.Println(createPointer())      // prints address of the memory location
	myFloatPointer := createPointer() // pointer is assigned to a variable
	fmt.Println(*myFloatPointer)      // получаем значение myFloat по адрессу myFloatPointer

	var myBool bool = true
	printPointer(&myBool) // pass the address of the variable for the pointer

	amount := 6
	Double(&amount)
	fmt.Println(amount)
	fmt.Println("====================================================")
	//todo=====================================================================

	//* Second case
	var value mySubscriber // creating a structure value
	value.old = 18
	var pointer *mySubscriber = &value // получаем указатель на значение структуры
	//Go посчитает, что поле myField должно содержать указатель. В действительности
	// это не так, поэтому происходит ошибка.
	//? fmt.Println(*pointer.old)
	//заключаем *pointer в круглые скобки, получаем значение mySubscriber, после
	//чего можем обратиться к полю структуры.
	fmt.Println((*pointer).old)
	//оператор «точка» позволяет обращаться к полям по указателям на структуры
	//точно так же, как вы обращаетесь к полям напрямую по значениям структур.
	//Круглые скобки и оператор * не обязательны.
	fmt.Println(pointer.old)

	applyDiscount(&value)
	fmt.Println(value.rate)

	value = *defaultSubcriber("Jon Ridik")
	//! Why value.old = 0
	fmt.Printf("%#v\n", value)
	valEasy := defaultSubcriber("Hren Gor'kiy")
	fmt.Printf("%#v\n", &valEasy)
	printInfo(valEasy)

}

// Объявляем, что функция возвращает указатель на float64.
func createPointer() *float64 {
	var myFloat = 99.66
	return &myFloat // возвращается адрес памяти переменной myFloat для указателя
}

// Указатели также можно передавать функциям в аргументах.
// Для этого достаточно указать, что один или несколько параметров имеют тип указателя.
func printPointer(myBoolPointer *bool) { // становится указателем для переменной myBool
	fmt.Println(*myBoolPointer) // выводит значение по адресу myBoolPointer
}

// func Double получает в аргументе адрес оригинального значения, а не копию,
// тем самым параметр становится указателем на это значение, а в самом коде
// с помощью оператора * мы получаем значение по адресу и изменяем его за
// пределами области функции Double
func Double(number *int) {
	*number *= 2
}

/*
Оператор & (амперсанд) используется в Go для получения адреса переменной.
Значения, представляющие адреса переменных, называются //* указателями
* Тип указателя
состоит из знака * и типа переменной, на которую ссылается указатель.
*int читается «указатель на int».
*/

//todo=====================================================================

// Изменяет поле rate существующей структуры за областью действия функции
func applyDiscount(s *mySubscriber) {
	s.rate = 4.99 // присваивание полю структуры по указателю без оператора *
}

// возвращает адрес оригинальных значений структуры для указателя, не забивая память.
func defaultSubcriber(name string) *mySubscriber {
	var s mySubscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}

// функции передаётся указатель, вызов полей происходит напрямую, оператор * ненужен.
func printInfo(s *mySubscriber) {
	fmt.Println("Name:", &s.name)
	fmt.Println("Monthly rate:", &s.rate)
	fmt.Println("Active?", &s.active)
}
