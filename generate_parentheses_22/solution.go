package generate_parentheses_22

func generateParenthesis(n int) []string {
	return genP(make([]string, 0), n, "")
}

func genP(strs []string, n int, s string) []string {
	if len(s) > n*2 {
		return strs
	} else if len(s) == n*2 {
		if validParens(s) {
			return []string{s}
		}
		return []string{}
	}

	arr1 := genP(strs, n, s+"(")
	arr2 := genP(strs, n, s+")")
	strs = append(strs, arr1...)
	strs = append(strs, arr2...)

	return strs
}

func validParens(s string) bool {
	stack := make([]string, len(s))
	top := 0
	for i := 0; i < len(s); i++ {
		if top > -1 && string(s[i]) == "(" {
			stack[top] = string(s[i])
			top++
		} else if string(s[i]) == ")" {
			top--
		}
	}

	return top == 0
}
