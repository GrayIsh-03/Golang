// func PaintSum import PaintNeeded <- paintneeded.go
package calculation

import "fmt"

// func определяет сколько нужно краски на каждую стену и сколько в общем
func PaintSum() {
	var total float64
	amount, err := PaintNeeded(12.2, 3)
	if err != nil {
		fmt.Println(err)
	}
	total += amount
	amount, err = PaintNeeded(8.3, -2.5)
	if err != nil {
		fmt.Println(err)
	}
	total += amount
	amount, err = PaintNeeded(12, 8)
	if err != nil {
		fmt.Println(err)
	}
	total += amount
	fmt.Printf("Total: need %0.2f a liters paint\n", total)
}
