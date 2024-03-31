package tasks

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	if x == 0 {
		return 0
	}
	num := x
	root := num / 2
	eps := 0
	for root-num/root > eps {
		root = int((root + num/root) / 2)
	}
	return root
}
