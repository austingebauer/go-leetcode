package move_zeroes_283

// Note: tricky problem!
// Ex: [1,0,0,3,12]
//            i
//        p
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
