package arrays

// Prefix and Suffix oriented task
func ProductExceptSelf(nums []int) []int {
	stateCompression := 1
	leftRotation := []int{}

	for i := 0; i < len(nums); i++ {
		leftRotation = append(leftRotation, stateCompression)
		stateCompression = stateCompression * nums[i]
	}

	rightRotation := make([]int, len(nums))
	stateCompression = 1

	for i := len(nums) - 1; i >= 0; i-- {
		rightRotation[i] = stateCompression
		stateCompression = stateCompression * nums[i]
	}

	result := []int{}

	for i := 0; i <= len(nums)-1; i++ {
		result = append(result, (rightRotation[i] * leftRotation[i]))
	}

	return result
}
