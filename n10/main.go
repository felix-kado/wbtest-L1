package main

import (
	"fmt"
	"math"
)

/*
Дана последовательность температурных колебаний:
-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
*/

func main() {
	// Типы обсуждаемы, точноть нам для такого примера явно не нужна, взял 64 потому что Mod такие использует
	seq := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[float64][]float64, len(seq))

	for _, item := range seq {
		groupInterval := item - math.Mod(float64(item), 10)
		groups[groupInterval] = append(groups[groupInterval], item)
	}

	fmt.Println(groups)
}
