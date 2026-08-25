package arrays

func RemoveElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 0

	for _, v := range nums {
		if v != val {
			nums[k] = v
			k++
		}
	}

	return k
}
