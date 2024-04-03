package tasks

//Мое решение не проходило по скорости, это чужое, проанализировать
func GetPal(s string, left int, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

func longestPalindrome(s string) string {
	var palindrom string
	for i, _ := range s {
		op1, op2 := GetPal(s, i, i), GetPal(s, i, i+1)
		if len(op1) > len(palindrom) {
			palindrom = op1
		}
		if len(op2) > len(palindrom) {
			palindrom = op2
		}
	}
	return palindrom
}
