package main

import (
	"bytes"
	"testing"
	"strings"
	"fmt"
)

const testFullResult = `|       ├───project
|               ├───file.txt (19b)
|               ├───gopher.png (70372b)
|       ├───static
|               ├───a_lorem
|                       ├───dolor.txt (empty)
|                       ├───gopher.png (70372b)
|                       ├───ipsum
|                               ├───gopher.png (70372b)    
|               ├───css
|                       ├───body.css (28b)
|               ├───empty.txt (empty)
|               ├───html
|                       ├───index.html (57b)
|               ├───js
|                       ├───site.js (10b)
|               ├───z_lorem
|                       ├───dolor.txt (empty)
|                       ├───gopher.png (70372b)
|                       ├───ipsum
|                               ├───gopher.png (70372b)    
|       ├───zline
|               ├───empty.txt (empty)
|               ├───lorem
|                       ├───dolor.txt (empty)
|                       ├───gopher.png (70372b)
|                       ├───ipsum
|                               ├───gopher.png (70372b)
|       ├───zzfile.txt (empty)
`

func TestTreeFull(t *testing.T) {
	out := new(bytes.Buffer)
	err := dirTree(out, "./testdata", 1)
	if err != nil {
		t.Errorf("test for OK Failed - error")
	}
    result := strings.TrimSpace(out.String())
    expected := strings.TrimSpace(testFullResult)
    if result != expected {
        t.Errorf("test for OK Failed - results not match\nGot:\n%v\nExpected:\n%v", result, expected)
    }
}