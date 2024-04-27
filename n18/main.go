package main

import (
	"fmt"
	"sync"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//  По завершению программа должна выводить итоговое значение счетчика.

type ConCounter struct {
	counter chan int
}

func NewCounter() ConCounter {
	c := ConCounter{counter: make(chan int, 1)}
	c.counter <- 0
	return c
}

func (c *ConCounter) increment() {
	cont := <-c.counter
	cont++
	c.counter <- cont
}

func (c *ConCounter) HowMuch() int {
	cont := <-c.counter
	res := cont
	c.counter <- cont
	return res
}

func main() {
	counter := NewCounter()
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.increment()
		}()
	}

	wg.Wait()
	fmt.Println(counter.HowMuch())

}
