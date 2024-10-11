package statistics

// Computes the slope and intercept of the linear regression line for datasets x and y.
func Regression(x, y []float64, xMean, yMean float64) (float64, float64) {
	var num, denom float64
	for i := 0; i < len(x); i++ {
		num += (x[i] - xMean) * (y[i] - yMean)
		denom += (x[i] - xMean) * (x[i] - xMean)
	}
	m := num / denom
	b := yMean - m*xMean
	return m, b
}