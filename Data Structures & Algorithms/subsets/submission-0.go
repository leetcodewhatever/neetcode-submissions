// in backtracking
// we are trying to answer 'What choice do I make right now, and how do I clean up after?'
// backtracking is Top Down.
func subsets(nums []int) [][]int {
	result := [][]int{}
	subset := []int{}
    
	var dfs func(nums []int, i int)
	dfs = func(nums []int, i int){
		if i == len(nums) {
			snapshot := make([]int, len(subset))
			copy(snapshot, subset)
			result = append(result, snapshot)
			return
		}
        
		// first choice: add a number
		subset = append(subset, nums[i])
		dfs(nums, i + 1)

		// second choice: don't add a number
		subset = subset[:len(subset)-1]
		dfs(nums, i + 1)

	}
	dfs(nums, 0)
	return result
}
