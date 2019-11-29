package main

import (
	"fmt"
	"sort"
)

func ExampleGrades_Descending() {
	gs := Grades{
		{
			name: "Tom",
			grade: 60,
		},
		{
			name: "John",
			grade: 100,
		},
	}

	sort.Sort(SortByDescending{gs})
	fmt.Println(gs)
	// Output:
	// [{Tom 60} {John 100}]
}

func ExampleGrades_Ascending() {
	gs := Grades{
		{
			name: "Tom",
			grade: 60,
		},
		{
			name: "John",
			grade: 100,
		},
	}

	sort.Sort(SortByAscending{gs})
	fmt.Println(gs)

	// Output:
	// [{John 100} {Tom 60}]
}
