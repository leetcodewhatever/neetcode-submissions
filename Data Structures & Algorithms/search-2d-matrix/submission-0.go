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
	
    var mid int
	mid = len(arr) / 2 
	for range len(arr) {

		if target == arr[mid] {
			return true
		}
        
	   if target > arr[mid] {
          mid = (len(arr) - mid) / 2
	   } else {
		  mid = mid / 2
	   }

	}
	return false
}