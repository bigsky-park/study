package main

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] == target {
		return left
	} else if nums[left] < target {
		return left + 1
	} else {
		return left
	}
}
