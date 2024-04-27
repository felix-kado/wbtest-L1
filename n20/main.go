package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func Reverse(s string) string {
	wordList := strings.Split(s, " ")

	for i := 0; i < len(wordList)/2; i++ {
		wordList[i], wordList[len(wordList)-i-1] = wordList[len(wordList)-i-1], wordList[i]
	}

	return strings.Join(wordList, " ")
}

func main() {

	fmt.Println("snow dog sun", "—", Reverse("snow dog sun"))
	fmt.Println("1 2 3 4 5 6 7 8 9", "—", Reverse("1 2 3 4 5 6 7 8 9"))
	fmt.Println("1 2 3 4 5 6 7 8", "—", Reverse("1 2 3 4 5 6 7 8"))
}
