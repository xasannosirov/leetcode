package solution

func sortColors(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		flag := true
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = false
			}
		}

		if flag {
			break
		}
	}
}
