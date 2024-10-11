package main

import (
	"bufio"
	"fmt"
	"guess-it/statistics"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers := []float64{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		numbs, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Fatalf("Error during conversion")
			return
		}
		numbers = append(numbers, numbs)

		if len(numbers) < 2 {
			// We need at least two data points to start making predictions
			fmt.Println("Please provide more numbers.")
			continue
		}

		// Perform linear regression and prediction
		x := make([]float64, len(numbers))
		for i := range x {
			x[i] = float64(i)
		}

		xMean := statistics.Average(x)
		yMean := statistics.Average(numbers)
		slope, intercept := statistics.Regression(x, numbers, xMean, yMean)

		// Predict the next number
		nextX := float64(len(numbers))
		predictedNext := slope*nextX + intercept

		// Calculate the range for the next number
		lower, upper := statistics.Range(numbers)

		fmt.Printf("%d %d %.2f\n", lower, upper, predictedNext)
	}
}
