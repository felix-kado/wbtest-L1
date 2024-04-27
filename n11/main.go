package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

// В обиходе ещё используется, множество не с пустой структурой, а с bool переменной.
// Мне сказали, что это не очень важно, у bool более лаконичная проверка наличия
// А пустая структура ничего не весит и не создает пространство неопредленности
// (Если у нас элемент с False лежит из-за ошибки какой-нибудь)

func intersection(a, b map[string]struct{}) map[string]struct{} {
	res := make(map[string]struct{})
	for el := range a {
		if _, exist := b[el]; exist {
			res[el] = struct{}{}
		}
	}
	return res
}

func main() {
	x := map[string]struct{}{
		"2s":     struct{}{},
		"3":      struct{}{},
		"513":    struct{}{},
		"asdasd": struct{}{},
	}
	y := map[string]struct{}{
		"2":     struct{}{},
		"3":     struct{}{},
		"513":   struct{}{},
		"Hello": struct{}{},
	}

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(intersection(x, y))
}
