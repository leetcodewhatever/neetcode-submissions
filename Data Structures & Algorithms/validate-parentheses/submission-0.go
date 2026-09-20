func isValid(s string) bool {
	var stack Stack[string]

	// map closing brackets to open brackets
	bracketsMap := make(map[string]string)
	bracketsMap[string('(')] = string(')')
	bracketsMap[string('[')] = string(']')
	bracketsMap[string('{')] = string('}')

	if len(s) == 1 {
		return false
	}
	
	s1 := string(s[0])
	for _, v := range bracketsMap {
		if v == s1 {
			return false
		}
	}

	for _, s := range s {
		// only push the opening brackets
		val := string(s)
		_, exists := bracketsMap[string(val)]

		if exists {
			stack.push(string(s))
		} else {
			if len(stack.items) != 0 && bracketsMap[stack.items[len(stack.items)-1]] == string(s) {
				stack.pop()
			} else {
              return false 
			} 
		}
	}

	if len(stack.items) == 0 {
		return true
	}

	return false

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
