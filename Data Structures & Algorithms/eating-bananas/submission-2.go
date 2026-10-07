import "slices"

func minEatingSpeed(piles []int, h int) int {
    low, high := 1, slices.Max(piles)
    // 1, 2 , 3
	for low < high { 
		mid := low + (high - low) / 2
		
		var hours int
		for _, p := range piles { 
		  // add mid - 1 to round up
          hours += (p + mid - 1) / mid
		}

		if hours > h {
			low  = mid + 1
		} else { 
			// if the current mid works, then we need to consider it an the values behind it.
			high = mid
		}
	}
	return low
}
