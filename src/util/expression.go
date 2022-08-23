package util

type Expression struct {
	Exp string
}

func (exp *Expression) GenSuffixExp() string {
	str := exp.Exp
	var res string
	var stack Stack[byte]
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			res += string(str[i])
		} else if stack.isEmpty() || str[i] == '(' {
			stack.Push(str[i])
		} else if str[i] == ')' {
			for stack.Top() != '(' {
				res += string(stack.Pop())
			}
			stack.Pop()
		} else {
			for getPrio(str[i]) <= getPrio(stack.Top()) {
				res += string(stack.Pop())
				if stack.isEmpty() {
					break
				}
			}
			stack.Push(str[i])
		}

	}
	for !stack.isEmpty() {
		res += string(stack.Pop())
	}
	return res
}
func (exp *Expression) Calculate() int {
	str := exp.Exp
	var s1 Stack[int]
	var s2 Stack[byte]
	number := 0
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			number = number*10 + int(str[i]-'0')
		} else if str[i] == '(' || s2.isEmpty() {
			if str[i] != '(' {
				s1.Push(number)
				number = 0
			}
			s2.Push(str[i])
		} else if str[i] == ')' {
			s1.Push(number)
			number = 0
			for s2.Top() != '(' {
				num1 := s1.Pop()
				num2 := s1.Pop()
				s1.Push(Operate(s2.Pop(), num2, num1))
			}
			s2.Pop()
		} else {
			for getPrio(s2.Top()) >= getPrio(str[i]) {
				num1 := s1.Pop()
				num2 := s1.Pop()
				s1.Push(Operate(s2.Pop(), num2, num1))
			}
			s2.Push(str[i])
		}
	}
	for !s2.isEmpty() {
		num1 := s1.Pop()
		num2 := s1.Pop()
		s1.Push(Operate(s2.Pop(), num2, num1))
	}
	return s1.Top()
}

func getPrio(ch byte) int {
	if ch == '*' || ch == '/' {
		return 2
	}
	if ch == '+' || ch == '-' {
		return 1
	}
	if ch == '(' {
		return 0
	}
	return -1
}

func Operate(ch byte, n1, n2 int) int {
	switch ch {
	case '+':
		return n1 + n2
	case '-':
		return n1 - n2
	case '*':
		return n1 * n2
	case '/':
		return n1 / n2
	}
	return -1

}
