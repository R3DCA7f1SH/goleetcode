package tasks

import "strings"

func lengthOfLastWord(s string) int {
	trimedString := strings.TrimSpace(s)
	sarray := strings.Split(trimedString, " ")
	return len(sarray[len(sarray)-1])
}
