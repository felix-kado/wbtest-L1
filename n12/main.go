package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func arrToSet(arr []string) map[string]struct{} {
	set := make(map[string]struct{})

	for _, item := range arr {
		if _, exist := set[item]; !exist {
			set[item] = struct{}{}
		}
	}

	return set
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(arrToSet(arr))
}
