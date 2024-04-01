package tasks

func runeToInt(rn rune) int {
	switch rn {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return ' '
	}

}

func romanToInt(s string) int {
	runesArray := []rune(s)
	length := len(runesArray)
	intSlice := []int{}
	for i := 0; i < length; i++ {
		intSlice = append(intSlice, runeToInt(runesArray[i]))
	}
	sum := 0
	for i := 0; i < length; i++ {
		if i < length-1 {
			if intSlice[i] >= intSlice[i+1] {
				sum += intSlice[i]
			} else {
				sum += intSlice[i+1] - intSlice[i]
				i++
			}
		} else {
			sum += intSlice[i]
		}
	}
	return sum
}
