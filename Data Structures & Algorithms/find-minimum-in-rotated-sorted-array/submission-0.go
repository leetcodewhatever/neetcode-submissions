func findMin(nums []int) int {
	/*
	   nums = [3,4,5,1,2]
	           l   m    r
	*/
	//
	//

	l := 0
	r := len(nums) - 1

	for l <= r {

		mid := l + (r-l)/2

		if nums[mid] > nums[r] {
			l = mid + 1
		} else {

			if nums[mid] < nums[l] {
				l = l + 1
			} else {
				r = mid - 1
			}

		}

	}

	return nums[l]

}