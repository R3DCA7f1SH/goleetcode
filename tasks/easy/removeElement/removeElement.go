package tasks

func removeElement(nums []int, val int) int {
	k := 0
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		if nums[0] == val {
			return 0
		}
		return 1
	}
	for i := 0; i < length; i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
