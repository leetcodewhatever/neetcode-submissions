func evalRPN(tokens []string) int {

    var stack Stack[int]

    for _, token := range tokens {

		if v, err := strconv.Atoi(token); err == nil {
		   stack.push(v)
		   continue
		}

		a := stack.pop()
		b := stack.pop()
		
		switch token {
			case "+":
				stack.push(a + b)
			case "-": 
				stack.push(b - a)
			case "*": 
				stack.push(a * b)
			case "/": 
				stack.push(b / a)
		}
	}
    
	return stack.pop()
}



type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) pop() T {
	lastIndex := len(s.items) - 1
	lastValue := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return lastValue 
}