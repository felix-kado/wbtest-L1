package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять
значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
Реализовать все возможные способы остановки выполнения горутины.
*/

// Просто генерация случайных данных
func getRandData() interface{} {
	time.Sleep(time.Millisecond * 50)
	switch rand.Intn(4) {
	case 0:
		return 1
	case 1:
		return "wow"
	case 2:
		return []int{1, 2, 3, 4}
	case 3:
		return map[int]string{1: "ф"}
	default:
		return nil
	}
}

func PublishToStream(ctx context.Context, ch chan<- interface{}) {
	// for select для выхода при отмене контекста
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefully ended publisher")
			return
		case ch <- getRandData():
		}
	}
}

func startWorker(ctx context.Context, ch <-chan interface{}) {
	// for select для выхода при отмене контекста
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefully ended worker")
			return
		case data := <-ch:
			fmt.Println(data)
		}
	}
}

func main() {
	// Количесвтво времени на работу программы будем брать из флага при запуске
	var nSecond int
	flag.IntVar(&nSecond, "n_second", 5, "Введите целое число")
	flag.Parse()

	if nSecond <= 0 {
		panic("n_sencond must be > 0 ")
	}
	// Писать будем сюда
	mainCh := make(chan interface{})
	// Для graceful shutdown создаем контекст отменяющийся через то количесво секунд, что мы указали при запуске
	ctx, stop := context.WithTimeout(context.Background(), time.Second*time.Duration(nSecond))
	defer stop()

	// Чтобы подождать все наши горутины
	var wg sync.WaitGroup

	// Создаем воркеры которые будут слушать канал и писать в stdout
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			startWorker(ctx, mainCh)
		}()
	}

	// Запускаем писателя
	wg.Add(1)
	go func() {
		PublishToStream(ctx, mainCh)
		wg.Done()
	}()

	// Блокируемся до отмены контекста
	<-ctx.Done()
	fmt.Println("Shutdown signal received")
	// Ждем пока все наши процессы как-то завершатся
	wg.Wait()

	close(mainCh)
	fmt.Println("All workers and publisher have exited")
}
