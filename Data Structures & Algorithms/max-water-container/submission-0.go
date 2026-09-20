func maxArea(height []int) int {
	/*


	  1. the chosen correct height is the smalest height
	  2. set r = left(nums) - 1
	  4. l = 1
	  5. height = min(height[l],hieght[r])
	  6. width = r - l
	  7. area = height * width
	  8.



	*/

	if len(height) == 2 {
		return min(height[0], height[1])
	}

	l := 0
	r := len(height) - 1
	var max_area int

	for {

		if l == len(height)-1 || r < 0 || l == r{
			break
		}

		h := min(height[l], height[r])
		w := r - l
		current_area := h * w
		max_area = max(max_area, current_area)

		if height[l] <= height[r] {
			l++
		}else{
			r--
		}

	}
	return max_area
}
