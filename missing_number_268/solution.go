package missing_number_268

func missingNumber(nums []int) int {
	n := len(nums)
	expectedSum := (n * (n + 1)) / 2
	actualSum := 0
	for _, n := range nums {
		actualSum += n
	}

	// the difference between expected and actual is the missing number
	return expectedSum - actualSum
}
