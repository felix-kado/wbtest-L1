package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

// Если я правильно понял последнее предложение, то это значит,
// что нужно получить строку, которая будут состоять из теж же рун, но в обратном порядке

func Reverse(s string) string {
	var builder strings.Builder
	// Заранее выделяем память, чтобы был минимум алокаций
	builder.Grow(len(s))

	for i := len(s); i > 0; {
		// Берем последнюю руну с конца среза
		r, size := utf8.DecodeLastRuneInString(s[:i])
		builder.WriteRune(r)
		// Уменьшаем длину среза
		i -= size
	}

	return builder.String()
}

func main() {
	// Правда у этой реализации есть забаный эффект, что составные символы распадаются, а диакритика может съехать))
	s1, s2, s3, s4 := "뢴", "🇩🇪", "möp", "главрыба"

	fmt.Printf("%s --- %s --- %s \n", s1, Reverse(s1), Reverse(Reverse(s1)))
	fmt.Printf("%s --- %s --- %s \n", s2, Reverse(s2), Reverse(Reverse(s2)))
	fmt.Printf("%s --- %s --- %s \n", s3, Reverse(s3), Reverse(Reverse(s3)))
	fmt.Printf("%s --- %s --- %s \n", s4, Reverse(s4), Reverse(Reverse(s4)))

}
