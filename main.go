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

		if len(numbers) > 1 {
			lower, upper := statistics.Range(numbers)
			fmt.Printf("%d %d\n", lower, upper)
		}
	}
}
