package main

import (
	"fmt"
	"math"
)

// Разработать программу, которая перемножает, делит,
// складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

func main() {
	// Интресно, неужели вся задача на то, чтобы тип int64 указать?
	var a, b int64

	a = int64(math.Pow(2, 27))
	b = int64(math.Pow(2, 37))

	sum := a + b
	diff := a - b
	product := a * b
	division := float64(a) / float64(b)

	fmt.Printf("%d + %d = %d\n", a, b, sum)
	fmt.Printf("%d - %d = %d\n", a, b, diff)
	fmt.Printf("%d * %d = %d\n", a, b, product)
	fmt.Printf("%d / %d = %f\n", a, b, division)

}
