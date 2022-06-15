package paintneed

import (
	"fmt"
)

func PaintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("Value a width =%0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("Value a height=%0.2f is invalid", height)
	}
	area := width * height
	fmt.Printf("Needed a %0.2f liters paint\n", area/10)
	return area / 10, nil
}
