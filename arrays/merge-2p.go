package arrays

// 88 merge sorted arrays
func Merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	if m == 0 {
		for i, v := range nums2 {
			nums1[i] = v
		}
		nums1 = nums2
		return
	}

	// two pointers for each array;
	firstPointer := m - 1
	secondPointer := n - 1

	// one pointer for positioning numbers in target array
	targetPointer := m + n - 1

	// two Pointers usually calls for checks ensuring both are inbounds

	// check first pointer is inbounds
	for secondPointer >= 0 {

		// check firstPointer is inbounds. if second pointer later is not inbounds anymore, we just add everything still valid with the other pointer
		if firstPointer >= 0 && nums1[firstPointer] > nums2[secondPointer] {

			// go from back to start to not override elements in nums1
			if nums1[firstPointer] != 0 {

				// swap elements to be safer
				nums1_temp := nums1[targetPointer]
				nums1[targetPointer] = nums1[firstPointer]
				nums1[firstPointer] = nums1_temp
			}

			firstPointer--

		} else {
			if nums2[secondPointer] != 0 {
				nums2_temp := nums1[targetPointer]
				nums1[targetPointer] = nums2[secondPointer]
				nums2[secondPointer] = nums2_temp
			}
			secondPointer--
		}

		targetPointer--
	}
}
