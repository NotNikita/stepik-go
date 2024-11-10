package main

import(
	"fmt"
	"strconv"
	"strings"
	// "sync"
)

// Dont need: select и context
/* // TODO:
1. Write linear code of each function
2. SingleHash
3. MultiHash
*/

// crc32(data)+"~"+crc32(md5(data))
func SingleHash(data int32) (result string) {
	str := strconv.Itoa(int(data))
	var arg1 = DataSignerCrc32(str)
	var arg2inner = DataSignerMd5(str)
	var arg2 = DataSignerCrc32(arg2inner)

	// Debug:
	result = arg1 + "~" + arg2
	fmt.Printf("%d SingleHash result %s\n", data, result)
	return
}

// th = [0;5]
// crc32(th+data)
func MultiHash(data string) (result string) {
	stack := make([]string, 0, 6)

	// routines?
	for th:=0; th<=5; th++ {
		thStr := strconv.Itoa(th)
		localResult := DataSignerCrc32(thStr + data)
		stack = append(stack, localResult)
	}
	result = strings.Join(stack, "")

	// Debug:
	fmt.Printf("%s MultiHash result: %s\n", data, result)
	return
}

func CombineResults(data []string) (result string) {
	result = strings.Join(data, "_")
	fmt.Printf("CombineResults \n %s", result)
	return
}

func main(){
	const ELEMENTS = 2
	results := make([]string, 0, ELEMENTS)

	for i:=0; i<ELEMENTS; i++ {
		i32 := int32(i)
		rez1 := SingleHash(i32)
		rez2 := MultiHash(rez1)
		results = append(results, rez2)
	}

	CombineResults(results)
}