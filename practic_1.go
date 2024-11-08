package main // week1

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func unique(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)
	var mapa = make(map[string]bool)

	for in.Scan() {
		txt := in.Text()
		if _, presented := mapa[txt]; !presented {
			mapa[txt] = true
			fmt.Fprintln(output, txt)
		} else {
			fmt.Errorf("Error: %s is already presented", txt)
		}
	}
	return nil
}

// cat practic_1_data.txt | go run practic_1.go
func main() {
	if err := unique(os.Stdin, os.Stdout); err != nil {
		panic(err)
	}

	fmt.Println("end")
}
