package tasks

//Решение не мое, проанализировать
func exist(board [][]byte, word string) bool {
	var backtrack func(int, int, int) bool
	backtrack = func(r, c int, i int) bool {
		if r < 0 || r > len(board)-1 || c < 0 || c > len(board[0])-1 || board[r][c] == '0' {
			return false
		}

		if board[r][c] != word[i] {
			return false
		}

		if i == len(word)-1 {
			return word[i] == board[r][c]
		}

		originalValue := board[r][c]

		board[r][c] = '0'

		if backtrack(r, c-1, i+1) || backtrack(r, c+1, i+1) || backtrack(r-1, c, i+1) || backtrack(r+1, c, i+1) {
			return true
		}

		board[r][c] = originalValue

		return false
	}

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			if backtrack(r, c, 0) {
				return true
			}
		}
	}

	return false
}
