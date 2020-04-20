package main

import (
	"fmt"
	"sort"
)

func main() {
	unSortedMap := map[string]int{"India": 20, "Canada": 70, "Germany": 15}

	// Int slice to store values of map.
	values := make([]int, 0, len(unSortedMap))

	for _, v := range unSortedMap {
		values = append(values, v)
	}

	// Sort slice values.
	sort.Ints(values)

	// Print values of sorted Slice.
	for _, v := range values {
		fmt.Println(v)
	}
}
