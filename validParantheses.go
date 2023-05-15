package main

type Stack []rune

func (s *Stack) Push(r rune) {
	*s = append(*s, r)
}

func (s *Stack) Pop() {
	n := len(*s)
	if n == 0 {
		return
	}
	(*s) = (*s)[:n-1]
}

func (s *Stack) Peek() rune {
	return (*s)[len(*s)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// question 20
func IsValid(str string) bool {
	if len(str) == 1 {
		return false
	}
	stack := Stack{}

	for _, s := range str {
		if s == '(' || s == '[' || s == '{' {
			stack.Push(s)
		}
		if stack.IsEmpty() {
			return false
		}
		if s == '}' {
			if stack.Peek() == '{' {
				stack.Pop()
			} else {
				return false
			}
		}
		if s == ')' {
			if stack.Peek() == '(' {
				stack.Pop()
			} else {
				return false
			}
		}
		if s == ']' {
			if stack.Peek() == '[' {
				stack.Pop()
			} else {
				return false
			}
		}
	}
	return stack.IsEmpty()
}
