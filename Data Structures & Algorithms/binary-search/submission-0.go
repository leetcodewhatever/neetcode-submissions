func search(nums []int, target int) int {
	//i := 0
	// start := 0
	// end := len(nums)
	// var index_tracker int
	// mid := left + (right-left)/2
	left := 0
	right := len(nums) - 1
	for left <= right {

		mid := left + (right-left) / 2

		if mid == len(nums)+1 {
			break
		}
		

		if nums[mid] == target {
			return mid
		}

		if  nums[mid] < target {
			//   nums = nums[mid:len(nums)]
			//   start = mid
			//   end = len(nums)
			left = mid + 1
		}

		if nums[mid] > target {
			right = mid - 1
		}

	}

	return -1

}