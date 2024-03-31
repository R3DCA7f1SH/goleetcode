package tasks

func plusOne(digits []int) []int {
	result := digits
	length := len(result)
	toAdd := 1
	for i := length - 1; i >= -1; i-- {
		if toAdd == 0 {
			break
		}
		var sum = result[i] + toAdd
		result[i] = sum % 10
		toAdd = sum / 10
		if i == 0 {
			if toAdd != 0 {
				// increase array length by 1
				result = append([]int{toAdd}, result...)
			}
			break
		}
	}
	return result
}
