package main // week1

import (
	"bytes"
	"strings"
	"testing"
	"bufio"
)

const (
    testOk = `1
2
3
4
5`

    expectedTestOk = `1
2
3
4
5
`
)

// go test -v
func TestUnique(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)

	if err := unique(in, out); err != nil {
		t.Errorf("test for OK Failed")
	}

	if out.String() != expectedTestOk {
		t.Errorf("test for OK Failed - results not match\nGot:\n%v\nExpected:\n%v", out.String(), expectedTestOk)
	}
}
