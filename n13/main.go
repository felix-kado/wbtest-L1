package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {
	me, you := "🤡", "🗿"

	me, you = you, me

	fmt.Printf("me: %s\n", me)
	fmt.Printf("you: %s\n", you)

}
