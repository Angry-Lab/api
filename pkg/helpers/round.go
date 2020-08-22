package helpers

import "math"

func Round(num, p float64) float64 {
	out := math.Pow(10, p)
	return math.Round(num*out) / out
}
