package product_of_array_except_self_238

// Uses prefix and suffix products to calculate the answer.
func productExceptSelf(nums []int) []int {
	// lProduct[i] is equal to the product of all values to the left of i
	lAcc := 1
	lProducts := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		lProducts[i] = lAcc
		lAcc = lAcc * nums[i]
	}

	// rProduct[i] is equal to the product of all values to the right of i
	rAcc := 1
	rProducts := make([]int, len(nums))
	for j := len(nums) - 1; j > -1; j-- {
		rProducts[j] = rAcc
		rAcc = rAcc * nums[j]
	}

	// the product at k is equal to the product of all values to the left
	// and right of k. Each of the products is precomputed and stored in
	// the lProducts and rProducts arrays.
	products := make([]int, len(nums))
	for k := 0; k < len(nums); k++ {
		products[k] = lProducts[k] * rProducts[k]
	}

	return products
}
