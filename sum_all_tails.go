package arrays_and_slices

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}

		sums = append(sums, Sum(numbers[1:]))
	}
	return
}
