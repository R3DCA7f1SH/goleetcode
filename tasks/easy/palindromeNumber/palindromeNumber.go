package tasks

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	var reverseNumber int
	for x > reverseNumber {
		reverseNumber = reverseNumber*10 + x%10
		x /= 10
	}
	return x == reverseNumber || x == reverseNumber/10
}
