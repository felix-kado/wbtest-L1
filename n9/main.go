package main

import (
	"fmt"
	"sync"
)

/*
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
 после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	testArr := []int{1, 3, 4, 5, 1, 23, 13, 4, 13, 4, 6, 3, 4, 4, 12, 3, 76, 7, 123, 4, 13, 3, 36, 8}

	intChan := make(chan int, len(testArr))
	intSqChan := make(chan int, len(testArr))

	// Из массива в канал
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(testArr))
		for _, num := range testArr {
			intChan <- num
			wg.Done()
		}
		wg.Wait()
		close(intChan)
	}()
	// Из первого канала квадрат во второй канал
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(testArr))
		for num := range intChan {
			intSqChan <- num * num
			wg.Done()
		}
		wg.Wait()
		close(intSqChan)
	}()
	// Из второго канала во stdout
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range intSqChan {
			fmt.Println(num)
		}
	}()
	wg.Wait()
}
