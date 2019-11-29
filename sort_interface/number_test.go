package main

import (
	"fmt"
	"sort"
)

func ExampleNumber() {
	n := Number{
		1, 10, 9, 0, 22,
	}

	sort.Sort(n)

	fmt.Println(n)
	// Output:
	// [0 1 9 10 22]
}
