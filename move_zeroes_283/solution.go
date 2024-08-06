package move_zeroes_283

// Ex: [1,0,0,3,12]
//
//	    i
//	p
func moveZeroes(nums []int) {
	placement := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			hold := nums[i]
			nums[i] = nums[placement]
			nums[placement] = hold
			placement++
		}
	}
}

func moveZeroes2(nums []int) {
	i, j := 0, 1
	for j < len(nums) {
		if nums[i] != 0 {
			i++
		}
		if nums[i] == 0 && nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
}
