package tasks

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	root := x >> 1
	for root-x/root > 0 {
		root = (root + x/root) >> 1
	}
	return root
}
