func dailyTemperatures(temperatures []int) []int {
	var stack Stack[int]
	ans := make([]int, len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {

		for len(stack.items) > 0 && temperatures[i] >= temperatures[stack.items[len(stack.items)-1]] {
			stack.pop()
		}
		if len(stack.items) == 0 {
			stack.push(i)
		} else {
			ans[i] = stack.items[len(stack.items)-1] - i
			stack.push(i)
		}

	}
	return ans
}



type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) pop() {
	lastItem := len(s.items) - 1
	if len(s.items) == 0 {
		return
	}
	s.items = s.items[:lastItem]
}