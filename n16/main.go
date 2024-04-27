package main

import (
	"fmt"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func QuickSort(arr []int) {
	// Алгоритм рекурсивный, вот базовый случай
	if len(arr) < 2 {
		return
	}

	// Два указателя
	left, right := 0, len(arr)-1

	pivot := arr[right]

	// Двигаем элементы меньшие апорного в левую часть
	for i := range arr {
		if arr[i] < pivot {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}
	// Ставим апорный элемент после меньших
	arr[left], arr[right] = arr[right], arr[left]

	// Запускаем алгоритм для левой и правой части
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5, 12, 23, 123, 12, 31, 5, 5, 132, 2, 3, 314, 4, 1, 23, 1, 41, 24, 6, 5, 554, 234, 77, 6, 4, 324, 11, 234, 45, 22, 34, 3, 4, 34}
	QuickSort(arr)
	fmt.Println("Sorted array:", arr)
}
