// TODO: revisit this problem
func findDuplicate(nums []int) int {
    fast, slow := 1, 0

	for fast < len(nums)  {
	
		if nums[slow] == nums[fast] {
			return nums[slow]
		}

		slow += 1
		fast += 2
	}

	return 0
}
