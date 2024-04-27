package main

import (
	"fmt"
)

// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

func whatsType(item interface{}) {
	switch item.(type) {
	case int:
		fmt.Println("I am int")
	case string:
		fmt.Println("I am string")
	case bool:
		fmt.Println("I am bool")
	case chan int:
		fmt.Println("I am chan int")
	}
}

func main() {

	whatsType(make(chan int))
	whatsType("asdasd")
	whatsType(123)
	whatsType(true)
}
