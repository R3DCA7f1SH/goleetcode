package tasks

func removeDuplicates(nums []int) int {
	k := 1
	length := len(nums)
	if length <= 1 {
		return 1
	}
	for i := 1; i < length; i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
