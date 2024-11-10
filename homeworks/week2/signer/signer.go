package main

import(
	"fmt"
	"strconv"
	"strings"
	"sync"
	"sort"
)

// Dont need: select и context
/* // TODO:
1. Write linear code of each function +
2. SingleHash +
3. MultiHash +
4. CombineResults +
5. Transform each function to receive 2 arguments: in and out channels +
6. ExecutePipeline -
*/

// crc32(data)+"~"+crc32(md5(data))
func SingleHash(in, out chan interface{}) {
	for data := range in {
		str := strconv.Itoa(int(data))
		var arg1 = DataSignerCrc32(str)
		var arg2inner = DataSignerMd5(str)
		var arg2 = DataSignerCrc32(arg2inner)
	
		// Debug:
		result := arg1 + "~" + arg2
		fmt.Printf("%d SingleHash result %s\n", data, result)
		out <- result
	}
}

// th = [0;5]
// crc32(th+data)
func MultiHash(in, out chan interface{}) {
	for data := range in {
		// for mapa access
		mu := &sync.Mutex{}
		// for handling routines
		wg := sync.WaitGroup{}
		mapa := map[int]string{}
		// Convert data to string
        dataStr := data.(string)

		thSize := 6
		wg.Add(thSize)
		for th:=0; th<thSize; th++ {
			go func(mu *sync.Mutex, num int) {
				thStr := strconv.Itoa(num)
				localResult := DataSignerCrc32(thStr + dataStr)
				mu.Lock()
				mapa[num] = localResult
				mu.Unlock()
				wg.Done()
			}(mu, th)
		}

		wg.Wait()
		var result []string
		for j:=0; j<thSize; j++ {
			result = append(result, mapa[j])
		}
		joinedSortedStrings := strings.Join(result, "")
	
		// Debug:
		fmt.Printf("%s MultiHash result: %s\n", data, joinedSortedStrings)
		out <- joinedSortedStrings
	}
}

// join(results, "_")
func CombineResults(in, out chan interface{}) {
	stack := make([]string, 0, ELEMENTS)

	for data := range in {
		stack = append(stack, data.(string))
	}

	sort.Strings(stack)

	result := strings.Join(stack, "_")

	// Debug:
	fmt.Printf("CombineResults \n %s", result)
	out <- result
	close(out)
}

// Receieves and manipultates functions
// func ExecutePipeline(jobs ...[]job) {
// 	wg := sync.WaitGroup{}

// }

const (
	ELEMENTS = 2
	MAX_ELEMENTS = 100
)

func main(){
	in := make(chan interface{}, MAX_ELEMENTS)
	out := make(chan interface{})
	// testExpected := "1173136728138862632818075107442090076184424490584241521304_1696913515191343735512658979631549563179965036907783101867_27225454331033649287118297354036464389062965355426795162684_29568666068035183841425683795340791879727309630931025356555_3994492081516972096677631278379039212655368881548151736_4958044192186797981418233587017209679042592862002427381542_4958044192186797981418233587017209679042592862002427381542"
	testForTwo = "29568666068035183841425683795340791879727309630931025356555_4958044192186797981418233587017209679042592862002427381542"
	runResult := "NOT_SET"

	inputData := [ELEMENTS]int{0,1}

	functions := []job{
		job(func(in, out chan interface{}){
			// it is closed only after all data has been sent.
			defer close(in)
			for _, num := range inputData {
				in <- num.(int32)
			}
		}),
		job(SingleHash),
		job(MultiHash),
		job(CombineResults),
		job(func(in, out chan interface{}) {
			dataRaw := <-in
			data, ok := dataRaw.(string)
			if !ok {
				t.Error("cant convert result data to string")
			}
			runResult = data
		}),
	}

	start := time.Now()

	ExecutePipeline(hashSignJobs...)

	end := time.Since(start)

	if testExpected != testResult {
		t.Errorf("results not match\nGot: %v\nExpected: %v", testResult, testExpected)
	}
}