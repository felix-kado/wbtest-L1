package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.

func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	var mid int

	for left <= right {
		mid = (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	arr := []int{1, 1, 1, 2, 3, 3, 4, 4, 4, 5, 5, 5, 5, 6, 6, 7, 8, 9, 10, 11, 12, 12, 22, 23, 23, 24, 31, 34, 34, 41, 45, 77, 123, 132, 234, 234, 314, 324, 554}
	target := 234
	index := BinarySearch(arr, target)
	fmt.Println(index)
}
