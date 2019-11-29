package main

type Number [] int

func (n Number) Len() int {
	return len(n)
}

func (n Number) Less(i, j int) bool {
	return n[i] < n[j]
}

func (n Number) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
