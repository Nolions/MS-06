package main

type Grade struct {
	name    string
	grade    float32
}

type Grades []Grade

//func (g Grade) Len() int
func (g Grades) Len() int {
	return len(g)
}

func (g Grades) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

type SortByAscending struct {
	Grades
}

func (s SortByAscending) Less(i, j int) bool {
	return s.Grades[i].grade > s.Grades[j].grade
}

type SortByDescending struct {
	Grades
}

func (s SortByDescending) Less(i, j int) bool {
	return s.Grades[i].grade < s.Grades[i].grade
}
