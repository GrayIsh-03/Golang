package main

import "fmt"

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
