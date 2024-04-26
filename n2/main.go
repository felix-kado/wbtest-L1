package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает
// значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func way1(nums []int) {
	// Ну чисто теоритически можно было бы напиать и без waitgroup и подождать, но не надо.
	var wg sync.WaitGroup
	for _, num := range nums {
		wg.Add(1)
		// В такой задаче то совсем не принципиально, но
		// вообще запускать отдельную горутину на вычесление
		// каджого квадрата это не очень хорошо, хоть они и не очень дорогие
		go func() {
			fmt.Println(num * num)
			defer wg.Done()
		}()
	}
	wg.Wait()

}

func way2(nums []int) {
	// Тут решение, которое тоже рассчитывает конкуретно,
	// но количество горутин не зависит от входных данных

	var wg sync.WaitGroup
	nWorkers := 8                   // количество горутин которые будут считать квадраты
	ch := make(chan int, len(nums)) // Буферизированный канал, из которого горутины смогут брать себе числа для вычисления квадратов

	// Наполняем канал
	for _, num := range nums {
		wg.Add(1)
		ch <- num
	}

	for i := 0; i < nWorkers; i++ {
		// Запускаем горутины воркеры
		go func() {
			for num := range ch {
				fmt.Println(num * num)
				wg.Done()
			}
		}()
	}
	wg.Wait()

}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	way1(nums)
	fmt.Println("-----")
	way2(nums)
}
