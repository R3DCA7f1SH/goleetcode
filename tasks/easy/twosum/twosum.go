package tasks

func twoSum(nums []int, target int) []int {
	currentMap := make(map[int]int)
	numsLength := len(nums)
	for i := 0; i < numsLength; i++ {
		_, currentOk := currentMap[nums[i]]
		_, anotherOk := currentMap[target-nums[i]]
		if anotherOk {
			return []int{currentMap[target-nums[i]], i}
		} else {
			if !currentOk {
				currentMap[nums[i]] = i
			}
		}
	}
	return []int{}
}
