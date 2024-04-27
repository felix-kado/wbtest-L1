package main

import (
	"fmt"
	"strings"
)

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

// У нас justString это глоабльная переменная, если мы просто полодим туда срез большой строки,
// полученный срез все еще будет ссылаться на исходный массив байтов, который содержит всю большую строку.
// И из-за этого сборщик мусора не сможет вычистить всю ненужную большую строку, поскольку будет активная ссылка.

var justString string

func createHugeString(size int) string {
	var builder strings.Builder

	builder.Grow(size)
	for i := 0; i < size; i++ {
		builder.WriteString("F")
	}
	return builder.String()
}

func someFunc() {
	v := createHugeString(1 << 10)
	temp := v[:100]

	var builder strings.Builder
	builder.Grow(len(temp))
	builder.WriteString(temp)

	justString = builder.String()
}

func main() {
	someFunc()
	fmt.Println(justString, len(justString))
}
