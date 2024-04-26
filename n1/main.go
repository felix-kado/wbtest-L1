package main

import "fmt"

// ЗАДАНИЕ
// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	height int
	Name   string
}

func (h *Human) Grow() {
	h.height += 1
}

type Action struct {
	Human    // Встраивание структуры Human, позволяет Action "наследовать" все поля и методы Human
	activity string
}

// AnnounceHeight выводит имя человека и его рост.
func (a Action) AnnounceHeight() {
	fmt.Printf("Привет, меня зовут %s, и мой рост %d см\n", a.Name, a.height) // Можем доставть атрибуты Human
}

func main() {
	person := Action{
		Human:    Human{Name: "Felix", height: 180}, // Инициализация структуры Human внутри Action
		activity: "бег трусцой",
	}

	person.AnnounceHeight()
	person.Grow() // Можем доставть методы Human
	person.AnnounceHeight()
}
