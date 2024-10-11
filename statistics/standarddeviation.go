package statistics

import "math"

// Calculates the standard deviation of a slice of float64 numbers.
func Deviation(num []float64) float64 {
	nb := Variance(num)
	numb := math.Sqrt(nb)
	return numb
}