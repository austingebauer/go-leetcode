package excel_sheet_column_number_171

import "math"

func titleToNumber(s string) int {
	var res int
	var exponent float64
	for i := len(s) - 1; i > -1; i-- {
		charNum := int(s[i]-'A') + 1
		res += charNum * int(math.Pow(float64(26), exponent))
		exponent++
	}
	return res
}
