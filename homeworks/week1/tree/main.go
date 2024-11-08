package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Notes:
/*
1. os.Chdir - change current directory
2. os.ReadDir - read directory
3. os.Readdir - Readdir reads the contents of the directory associated with file and returns a slice of up to n FileInfo values
4. os.Readdirnames

*/

func dirTree(out io.Writer, directory string, level int) error {
	var files, err = os.ReadDir(directory)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}

	for _, entry := range files {
		if entry.IsDir() {
			// print dir name and call recursive
			fmt.Fprintf(out, "|%s├───%s\n", strings.Repeat("\t", level), entry.Name())
			dirTree(out, directory+"/"+entry.Name(), level+1)
		} else {
			// its a file, print name and size
			fileInfo, err := entry.Info()
			if err != nil {
				fmt.Errorf("Error: %s", err)
			}
			var sizeInfo = ""
			if fileInfo.Size() == 0 {
				sizeInfo = "empty"
			} else {
				sizeInfo = strconv.FormatInt(fileInfo.Size(), 10) + "b"
			}
			fmt.Fprintf(out, "|%s├───%s (%s)\n", strings.Repeat("\t", level), fileInfo.Name(), sizeInfo)
		}
	}
	return nil
}

func main() {
	out := new(bytes.Buffer)
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	err := dirTree(out, "./", 1)
	// fmt.Println(out.String())
	if err != nil {
		panic(err.Error())
	}
}
