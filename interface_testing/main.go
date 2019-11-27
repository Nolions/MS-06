package main

import "fmt"

func main() {
	c := new(chinese)
	c.Name = "sam"

	fmt.Println(sayHi(c))
	fmt.Println(sayHello(c))

}

func sayHi(g Greet) string {
	return fmt.Sprintf("%s %s", g.Hi(), g.GetName())
}

func sayHello(g Greet) string  {
	return fmt.Sprintf("%s %s", g.Hello(), g.GetName())
}

