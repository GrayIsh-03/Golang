package hello

import (
	"rsc.io/quote"
)

func Hello() string {
	return quote.Hello()
}

func double(n float64) float64 {
	n *= n
	return n
}
