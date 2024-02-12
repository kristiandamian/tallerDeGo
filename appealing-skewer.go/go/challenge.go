package main

import (
	"fmt"
)

func lastNode(len int) int {
	v := 1
	for i := len; i > 1; i-- {
		v = v * i
	}

	return v
}

func numberOfNodes(len int) int {
	total := 1
	for p := 1; p < len; p++ {
		prev := 1
		for i := len; i > (len - p); i-- {
			prev = prev * i
		}
		total += prev
	}

	return total
}

func duplicateSlice[T any](src []T) []T {
	dup := make([]T, len(src))
	copy(dup, src)
	return dup
}

func permute(value []string, pos int, c chan []string) {
	if pos <= len(value) {
		for i := pos; i < len(value); i++ {
			current := value[pos]
			new := value[i]
			newValue := duplicateSlice(value)
			newValue[pos] = new
			newValue[i] = current

			if pos == len(value)-1 {
				c <- newValue
			} else {
				newPos := pos + 1
				go permute(newValue, newPos, c)
			}
		}
	}
}

func getAllPosibilities(value []string) [][]string {
	var posibilities = [][]string{}
	lastNodeLength := lastNode(len(value))
	c := make(chan []string, lastNodeLength)

	if len(value) == 0 {
		return nil
	}

	go permute(value, 0, c)

	for i := 0; i < lastNodeLength; i++ {
		posibilities = append(posibilities, <-c)
	}

	return posibilities
}

func calculateValidCombinations(ingredients []string) int {
	// write your code here -----------------------------
	if len(ingredients) == 0 {
		return 0
	}
	total := 0

	for _, v := range getAllPosibilities(ingredients) {
		dup := false
		for i := 0; i < len(v)-1; i++ {
			if v[i] == v[i+1] {
				dup = true
				break
			}
		}
		if !dup {
			total++
		}
	}

	return total
	// --------------------------------------------------
}

func main() {
	ingredients := []string{"🥔", "🥩", "🥩"}
	expected := 2

	actual := calculateValidCombinations(ingredients)

	if expected == actual {
		fmt.Printf("SUCCESS | expected: %d | got: %d", expected, actual)
	} else {
		fmt.Printf("FAILED | expected: %d | got: %d", expected, actual)
	}
}
