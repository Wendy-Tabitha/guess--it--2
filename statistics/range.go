package statistics

func Range(numbers []float64) (int, int) {
	x := make([]float64, len(numbers))
	for i := range x {
		x[i] = float64(i)
	}

	xMean := Average(x)
	yMean := Average(numbers)
	slope, intercept := Regression(x, numbers, xMean, yMean)

	nextX := float64(len(numbers))
	predictedNext := slope*nextX + intercept

	// avg := Average(numbers)
	stdDev := Deviation(numbers)

	// Adjust factor as needed based on your data
	factor := 2.57

	lower := int(predictedNext - factor*stdDev)
	upper := int(predictedNext + factor*stdDev)
	return lower, upper
}
