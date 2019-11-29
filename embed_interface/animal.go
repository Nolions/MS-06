package main

import "sort"

type Animal interface {
	Speak() string
	fly() bool
}

type Dog struct{
	Animal
}

func (d *Dog) Speak() string {
	return "Wow"
	sort.Float64s()
}

//func (d *Dog) fly() bool {
//	return false
//}