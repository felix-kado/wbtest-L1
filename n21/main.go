package main

import "fmt"

// Реализовать паттерн «адаптер» на любом примере

// Humanoid -- интерфейс определяет поведение для гуманоидных существ
type Humanoid interface {
	Speak() string
	Walk() string
}

// Структура, реализующая интерфейс Humanoid
type Human struct{}

// Позволяет человеку говорить
func (h Human) Speak() string {
	return "Привет, я человек"
}

// Позволяет человеку ходить
func (h Human) Walk() string {
	return "Иду"
}

// Это структура со своими методами, отличающимися от Humanoid
type Alien struct{}

// Издает звуки, характерные для инопланетян
func (a Alien) EmitSound() string {
	return "Блорг! Зоп Мип зорп"
}

// Осуществляет перемещение инопланетянина
func (a Alien) Move() string {
	return "Леверирую"
}

// Обеспечивает адаптацию Alien к интерфейсу Humanoid
type AlienAdapter struct {
	alien Alien
}

// Адаптирует метод EmitSound инопланетянина к методу Speak гуманоида
func (aa AlienAdapter) Speak() string {
	return aa.alien.EmitSound()
}

// Адаптирует метод Move инопланетянина к методу Walk гуманоида
func (aa AlienAdapter) Walk() string {
	return aa.alien.Move()
}

// Просто для теста
func humanoidActivities(h Humanoid) {
	fmt.Println(h.Speak())
	fmt.Println(h.Walk())
}

func main() {
	human := Human{}
	alien := Alien{}
	alienAdapter := AlienAdapter{alien: alien}

	fmt.Println("Действия человека:")
	humanoidActivities(human)

	fmt.Println("Действия инопланетянина через адаптер:")
	humanoidActivities(alienAdapter)
}
