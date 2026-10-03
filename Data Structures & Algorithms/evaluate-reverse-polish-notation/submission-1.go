func evalRPN(tokens []string) int {

    var stack Stack[int]

    for _, token := range tokens {
		switch token {
			case "+":
				a := stack.pop()
				b := stack.pop()
				stack.push(a+b)
			case "-": 
			    a := stack.pop()
				b := stack.pop()
				stack.push(b - a)
			case "*": 
			    a := stack.pop()
				b := stack.pop()
				stack.push(a*b)
			case "/": 
			    a := stack.pop()
				b := stack.pop()
				stack.push(a/b)
			default:
				v, _ := strconv.Atoi(token)

				if len(tokens) == 1 {
					return v
				}

				stack.push(v)

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