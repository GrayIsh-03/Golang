package calculation

import "fmt"

//func PaintNeeded рассчитывает сколько потребуется краски на стену определённой
//ширины и высоты, с учётом расхода 1 литр на 10 кв.м.
func PaintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("value a width =%0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("value a height=%0.2f is invalid", height)
	}
	area := width * height
	fmt.Printf("Needed a %0.2f liters paint\n", area/10)
	return area / 10, nil
}
