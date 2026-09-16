package main

import (
	"leetcode/arrays"
)

// Vergleiche die größten noch nicht verarbeiteten Werte und befülle nums1 von hinten.
// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively. Merge nums1 and nums2 into a single array sorted in non-decreasing order. The final sorted array should not be returned by the function, but instead be *stored inside the array *nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

func main() {
	callHIndex()
}

func callHIndex() {
	array1 := []int{3, 0, 6, 1, 5}
	arrays.HIndexBucketVersion(array1)
}

func callCanJump() {
	array1 := []int{1, 0, 1, 0}
	arrays.CanJump(array1)

	array1 = []int{2, 3, 1, 1, 4}
	arrays.CanJump(array1)
}

func callCanJumpII() {
	array1 := []int{2, 3, 1, 1, 4}
	arrays.Jump(array1)
}

func callMaxProfit() {
	// array1 := []int{7, 6, 4, 3, 1}
	// arrays.MaxProfit(array1)
	array1 := []int{1, 2}
	arrays.MaxProfit(array1)
}

func callRotateShiftRight() {
	array1 := []int{1, 2, 3, 4, 5, 6, 7}
	arrays.RotateShiftRight(array1, 3)
}

func callRemoveDuplicatesFromSortedArrayII() {
	array1 := []int{1, 1, 1, 2, 2, 3}
	arrays.RemoveDuplicatesFromSortedArrayII(array1)
}

func callRemoveDuplicatesFromSortedArray() {
	array1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	arrays.RemoveDuplicatesFromSortedArray(array1)

	array1 = []int{2, 2, 3}
	arrays.RemoveDuplicatesFromSortedArray(array1)
}

func callRemoveElements() {
	array1 := []int{3, 2, 2, 3}
	val := 3

	arrays.RemoveElement(array1, val)
}

func callArrayMerge() {
	array1 := []int{0, 0, 3, 0, 0, 0, 0, 0, 0}
	array2 := []int{-1, 1, 1, 1, 2, 3}
	arrays.Merge(array1, 3, array2, 6)

	array1 = []int{1}
	array2 = []int{}
	arrays.Merge(array1, 1, array2, 0)
}
