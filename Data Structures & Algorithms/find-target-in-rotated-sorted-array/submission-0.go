func search(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

 
	for left < right {
       
		mid := left + (right-left) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			if target <= nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}

		}else{
			if target >= nums[mid] && target <= nums[right] {
				left = left + 1

			}else{
				right = right - 1
			}
		}

	}

    if nums[left] != target{
        return -1
    }else{
        return left
    }

}