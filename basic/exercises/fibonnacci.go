package exercises

func calculate(position int, elements []int) int {
	if position <= 1 {
		return 1
	} else {
		return elements[position-1] + elements[position-2]
	}
}

func Fibonnacci(numerOfelements int) []int {
	var elements = []int{}
	for i := 0; i < numerOfelements; i++ {
		elements = append(elements, calculate(i, elements))
	}

	return elements
}
