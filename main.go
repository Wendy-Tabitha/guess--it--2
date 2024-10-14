package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"guess-it/statistics"
)

func main() {
	numbers := []float64{}

	inputSize := 1

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		numbs, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Fatalf("Error during conversion")
			return
		}
		numbers = append(numbers, numbs)

		if len(numbers) == 1 {
			// We need at least two data points to start making predictions
			fmt.Println("Please provide more numbers.")
			continue
		}

		allSame := true
		firstValue := numbers[0]
		for _, value := range numbers {
			if value != firstValue {
				allSame = false
				break
			}
		}

		if allSame {
			fmt.Println("All values in the dataset are identical")
			return
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
		nextX := float64(inputSize)
		predictedNext := slope*nextX + intercept

		// Calculate the range for the next number
		lower, upper := statistics.Range(numbers)

		fmt.Printf("%d %d %.2f\n", lower, upper, predictedNext)
	}
}
