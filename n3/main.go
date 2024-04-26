package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func way1(nums []int) int {
	// Ну чисто теоритически можно было бы напиать и без waitgroup и подождать, но это совсем странное.
	var wg sync.WaitGroup
	resCh := make(chan int, len(nums)) // Буферизированный канал, из которого горутины смогут брать себе числа для вычисления квадратов

	for _, num := range nums {
		wg.Add(1)
		// В такой задаче то совсем не принципиально, но
		// вообще запускать отдельную горутину на вычесление
		// каджого квадрата это не очень хорошо, хоть они и не очень дорогие
		go func() {
			resCh <- num * num
			defer wg.Done()
		}()
	}
	wg.Wait()
	close(resCh)

	summ := 0
	for i := range resCh {
		summ += i
	}

	return summ
}

func way2(nums []int) int {
	// Тут решение, которое тоже рассчитывает конкуретно,
	// но количество горутин не зависит от входных данных

	var wg sync.WaitGroup
	nWorkers := 8                   // количество горутин которые будут считать квадраты
	ch := make(chan int, len(nums)) // Буферизированный канал, из которого горутины смогут брать себе числа для вычисления квадратов
	resCh := make(chan int, 1)
	resCh <- 0
	// Наполняем канал
	for _, num := range nums {
		wg.Add(1)
		ch <- num
	}

	for i := 0; i < nWorkers; i++ {
		// Запускаем горутины воркеры
		go func() {
			for num := range ch {
				summ := <-resCh
				summ += num * num
				resCh <- summ
				wg.Done()
			}
		}()
	}
	wg.Wait()

	return <-resCh
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	fmt.Println(way1(nums))
	fmt.Println("-----")
	fmt.Println(way2(nums))
}
