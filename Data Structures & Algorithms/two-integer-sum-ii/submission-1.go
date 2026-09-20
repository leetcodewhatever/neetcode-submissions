func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1
    var result []int
	for l != r {
	  v := numbers[l] + numbers[r]
	  if v == target{
		result = append(result, l+1) // add extra one, since we are assuming it's a 1 index array.
		result = append(result, r+1)
		return result
	  }
	  if v > target{ r = r - 1 }
	  if v < target { l = l + 1 }
	}

return result

}
