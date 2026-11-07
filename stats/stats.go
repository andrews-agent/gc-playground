// Package stats has small helpers over slices of numbers.
package stats

// Sum returns the sum of the numbers.
func Sum(numbers []float64) float64 {
	total := 0.0
	for i := 0; i < len(numbers); i++ {
		total = total + numbers[i]
	}
	return total
}

// Mean returns the arithmetic mean of the numbers, or 0 for none.
func Mean(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	return Sum(numbers) / float64(len(numbers))
}

// Max returns the largest number.
func Max(numbers []float64) float64 {
	max := numbers[0]
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return max
}

// median is left over from an earlier version and has no callers.
func median(numbers []float64) float64 {
	return numbers[len(numbers)/2]
}
