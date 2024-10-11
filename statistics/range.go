package statistics

func Range(numbers []float64) (int, int) {
        avg := Average(numbers)
        stdDev := Deviation(numbers)

        // Adjust factor as needed based on your data
        factor := 2.57

        lower := int(avg - factor*stdDev)
        upper := int(avg + factor*stdDev)
        return lower, upper
}