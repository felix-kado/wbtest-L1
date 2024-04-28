package main

import (
	"context"
	"fmt"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

// Остановка горутины с использованием канала
func goroutineWithChannel(stopChan chan bool) {
	for {
		select {
		case <-stopChan:
			fmt.Println("Горутина остановлена через канал")
			return
		default:
			fmt.Println("Работает горутина с каналом")
			time.Sleep(1 * time.Second)
		}
	}
}

// Остановка горутины с использованием контекста
func goroutineWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина остановлена через контекст")
			return
		default:
			fmt.Println("Работает горутина с контекстом")
			time.Sleep(1 * time.Second)
		}
	}
}

func goroutine() {
	fmt.Println("Горутина просто закончила свою работу")
}

func panicWorker() {
	fmt.Println("Горутина начала работу")
	time.Sleep(1 * time.Second) // Имитация работы
	panic("Произошла паника в горутине")
}

func main() {
	// Использование канала для остановки горутины
	stopChan := make(chan bool)
	go goroutineWithChannel(stopChan)
	time.Sleep(2 * time.Second)
	stopChan <- true

	// Использование контекста для остановки горутины
	ctx, cancel := context.WithCancel(context.Background())
	go goroutineWithContext(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

	// Горутина может просто закончить работу
	go goroutine()
	time.Sleep(1 * time.Second)

	// Может паника случиться
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Перехвачена паника:", r)
			}
		}()
		panicWorker()
	}()

	time.Sleep(2 * time.Second)
}
