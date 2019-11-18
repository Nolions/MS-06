package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Age   int    `json:age`
}

func main() {
	u := user{
		Name:  "Sam",
		Email: "123@123.123",
		Phone: "123409865",
		Age:   20,
	}

	s := encode(u)
	fmt.Println(s)

	d := decode(s)
	fmt.Println(d)
}

// json encode
func encode(u user) string {
	d, err := json.Marshal(u)
	if err != nil {
		log.Fatal("error:", err)
	}
	return string(d)
}

// json decode
func decode(d string) user {
	u := user{}
	json.Unmarshal([]byte(d), &u)

	return u
}
