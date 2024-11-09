package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const (
	iterations = 7
	goRoutines = 5
)

func doWork (routineI int) {
	for j := 0; j < iterations; j++ {
		fmt.Println(formatWork(routineI, j))
		runtime.Gosched()
	}
}

func main() {
	for i:=0; i < goRoutines; i++ {
		go doWork(i)
	}
	fmt.Scanln()
}

func formatWork (routineI int, iter int) string {
	return fmt.Sprintln(strings.Repeat("  ", routineI), "█",
		strings.Repeat("  ", goRoutines - routineI),
		"th", routineI,
		"iter", iter, strings.Repeat("■", iter))
}

func imports() {
	fmt.Println(time.Millisecond, runtime.NumCPU())
}
