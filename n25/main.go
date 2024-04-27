package main

import (
	"fmt"
	"time"
)

func MySleep(sec int) {
	// Создаем таймер, который отправляет значение в канал после заданного времени
	timer := time.NewTimer(time.Duration(sec) * time.Second)
	<-timer.C // Ожидаем сигнал из канала таймера, это блоирующая операция
}

func main() {
	fmt.Println("Город засыпает")
	MySleep(2)
	fmt.Println("Город просыпается, никого не убили")
}
