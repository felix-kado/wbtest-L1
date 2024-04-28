package main

import (
	"context"
	"fmt"
)

// Реализовать все возможные способы остановки выполнения горутины.

// Пока не делал!
func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped")
			return
		default:

		}
	}
}

func main() {

}
