package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.

func main() {
	var wg sync.WaitGroup

	// Мапа с обычным мютексом, т.к без него будет гонка или паника при одновременной записи
	mapa := make(map[int]string)
	var mu sync.Mutex
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()

			mapa[i] = "done"
		}()
	}
	wg.Wait()
	fmt.Println(mapa)

	// Ну или можно с помощью sync.Map там уже чтение и запись атомарные

	mapa2 := new(sync.Map)
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mapa2.Store(i, "done")
		}()
	}
	wg.Wait()
	fmt.Println(mapa)
}
