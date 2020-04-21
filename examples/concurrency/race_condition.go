package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wait sync.WaitGroup
var count int

func increment(s string) {
	for i := 0; i < 10; i++ {
		x := count
		x++
		time.Sleep(time.Duration(rand.Intn(4)) * time.Millisecond)
		count = x
		fmt.Println(s, i, "Count: ", count)

	}
	wait.Done()

}
func main() {
	wait.Add(2)
	go increment("foo: ")
	go increment("bar: ")
	wait.Wait()
	fmt.Println("last count value ", count)
}

// The count resource is accessed by 2 go routines.
// Each routine iterates to 10 times.
// In such case, the count variable should be 20 at last.
// But it is not so because it is simulating race condition.

// Output:

// foo:  0 Count:  1
// bar:  0 Count:  1
// foo:  1 Count:  2
// foo:  2 Count:  3
// foo:  3 Count:  4
// bar:  1 Count:  2
// foo:  4 Count:  5
// foo:  5 Count:  6
// foo:  6 Count:  7
// bar:  2 Count:  3
// bar:  3 Count:  4
// bar:  4 Count:  5
// foo:  7 Count:  8
// foo:  8 Count:  9
// bar:  5 Count:  6
// bar:  6 Count:  7
// foo:  9 Count:  10
// bar:  7 Count:  8
// bar:  8 Count:  9
// bar:  9 Count:  10
// last count value  10
