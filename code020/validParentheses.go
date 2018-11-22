package code020

type stack []rune

var matching = map[rune]rune{
	')':'(',
	']':'[',
	'}':'{',
}

func isValidParenteses(input string) bool {
	s := new(stack)

	for _, b := range input  {
		switch b {
		case '(', '[', '{':
			s.push(b)
		case ')', ']', '}':
			if r, ok := s.pop(); !ok || r != matching[b] {
				return false
			}
		}
	}

	if len(*s) > 0 {
		return false
	}
	return true
}

func(s *stack)push(b rune) {
	*s = append(*s, b)
}

func(s *stack)pop() (rune, bool) {
	if len(*s) <= 0 {
		return 0, false
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res, true
}
