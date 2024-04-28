package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
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
			close(ch)
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
	// Количесвтво воркеров будем брать из флага при запуске
	var nWorkers int
	flag.IntVar(&nWorkers, "n_workers", 4, "Введите целое число")
	flag.Parse()

	// Писать будем сюда
	mainCh := make(chan interface{})

	// Для graceful shutdown создаем контекст отменяющийся по сигналу Interrupt нашему процессу
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Чтобы подождать все наши горутины
	var wg sync.WaitGroup

	// Создаем воркеры которые будут слушать канал и писать в stdout
	for i := 0; i < nWorkers; i++ {
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

	fmt.Println("All workers and publisher have exited")
}
