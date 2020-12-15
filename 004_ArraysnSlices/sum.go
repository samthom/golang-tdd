package arraynslices

// Sum add slice
func Sum(numbers []int) (sum int) {
	for _, i := range numbers {
		sum += i
	}
	return
}

// SumAll variable sum
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}
