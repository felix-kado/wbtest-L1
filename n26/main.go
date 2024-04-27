package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

func isUnique(s string) bool {
	charSet := make(map[rune]struct{})
	s = strings.ToLower(s)

	for _, char := range s {
		if _, exists := charSet[char]; exists {
			return false
		}
		charSet[char] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println(isUnique("abcd"))
	fmt.Println(isUnique("abCdefAaf"))
	fmt.Println(isUnique("aabcd"))
}
