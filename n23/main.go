package main

import "fmt"

// Удалить i-ый элемент из слайса.

func removeElementInplace1(slice []int, i int) []int {
	slow := 0

	for fast := 0; fast < len(slice); fast++ {
		if fast == i {
			fast++
		}
		slice[slow] = slice[fast]
		slow++
	}
	return slice[:slow]
}
func removeElementInplace2(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func removeElementInplace3(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(nums)
	fmt.Println(removeElementInplace1(nums, 3))
	fmt.Println(nums)
	fmt.Println(removeElementInplace2(nums, 3))
	fmt.Println(nums)
	fmt.Println(removeElementInplace3(nums, 3))
	fmt.Println(nums)

}
