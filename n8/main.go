package main

import "fmt"

func setBit(num int64, i uint) int64 {
	// Побитовое ИЛИ с масокой 1 в i разряде, остальные 0
	return num | (1 << i)
}

func clearBit(num int64, i uint) int64 {
	// Побитовое И c масокой 0 в i разряде, остальные 1
	return num & ^(1 << i)
}

func main() {
	var num int64 = 0
	var index uint = 3

	num = setBit(num, index)
	fmt.Println(num)

	num = clearBit(num, index)
	fmt.Println(num)

}
