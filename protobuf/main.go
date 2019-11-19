package main

import (
	"MS-06/protobuf/protobuf"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	user := protobuf.User{
		Id:       1,
		Username: "sam",
	}

	d := decode(&user)
	fmt.Println(d)

	fmt.Println(encode(d))
}

// 編碼成Protocol Buffer
func decode(user *protobuf.User) []byte {
	data, err := proto.Marshal(user)
	if err != nil {
		log.Fatal(err.Error())
	}

	return data
}

// Protocol Buffer 解碼
func encode(d []byte) *protobuf.User {
	u := protobuf.User{}
	proto.Unmarshal(d, &u)

	return &u
}
