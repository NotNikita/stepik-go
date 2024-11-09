package main

import (
	"fmt"
)


func main() {
	// bufferized channel
	ch := make(chan int, 1)

	// creating go routine
	go func(in chan int) {
		val := <- in
		fmt.Println("Go get from chan", val)
		fmt.Println("Go after reading")
	}(ch)

	ch <- 10
	ch <- 20

	fmt.Println("Main function finished")
	fmt.Scanln()
}
// main -> get -> after get