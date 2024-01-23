package exercises

func sum(c chan int) {
	x := <-c
	y := <-c
	c <- x + y
}

func FirstRoutine() int {
	c := make(chan int)
	go sum(c)

	c <- 10
	c <- 15

	r := <-c

	return r
}
