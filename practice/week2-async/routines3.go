package main

import (
	"fmt"
)


func main() {
	// bufferized channel
	ch := make(chan int)

	// creating go routine, that accepts channel for writing
	go func(out chan<- int) {
		for i := 0; i<4; i++ {
			fmt.Println("before", i)
			out <- i
			fmt.Println("after", i)
		}
		close(out) // if this line removed: fatal error: all goroutines are asleep - deadlock!

		fmt.Println("Generator finished")
	}(ch)

	// Конструкция close помогает завершить цикл, итерирующийся по значениям, идущим в канал
	for i := range ch {
		fmt.Println("\tget", i)

	}
	// fmt.Scanln()
}
/*
before 0
after 0
before 1
get 0
get 1
after 1
before 2
after 2
before 3
get 2
get 3
after 3
Generator finished
*/