func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix[0])
	for r := range matrix {
		// fmt.Println(r)
		for range matrix[0] {

			if target > matrix[r][n - 1] {
				break
			}

			if target < matrix[r][0] {
				break
			}

			if binarySearch(matrix[r], target) {
				return true
			}
		}
	}
    return false
}

func binarySearch(arr []int, target int) bool {
	
    low, high := 0, len(arr) - 1
	for low <= high {

		mid := low + (high - low) / 2

		if target == arr[mid] {
			return true
		}
        
	   if target > arr[mid] {
		  low = mid + 1
	   } else {
		  high = mid - 1
	   }

	}
	return false
}