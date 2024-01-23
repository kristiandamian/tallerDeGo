package main

import (
	"fmt"
	"tests/exercises"
)

func main() {
	fmt.Println(exercises.Fibonnacci(10))
	fmt.Println("Es palindro", exercises.IsPalindrom("Anita lava la tina"))
	fmt.Println("Es palindro", exercises.IsPalindrom("mock"))
	fmt.Println("Go Routine", exercises.FirstRoutine())
}
