package main

import "fmt"

func main() {
	a := new(Dog)
	Say(a)
}

func Say(a Animal) {
	fmt.Println(a.Speak())
}
