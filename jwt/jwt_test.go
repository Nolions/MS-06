package jwt

import "fmt"

func ExampleSigning() {
	jwtToken := signing(&user{
		id:      2,
		account: "sam78",
	})
	_, err := validating(jwtToken)
	fmt.Println()

	// output:
	// true
}

func ExampleValidating_Fail() {
	_, err := validating("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50Ijoic2FtNzgiLCJpZCI6MX0.2pWmbRvrQ3FvEorhJEtwu3POOaO8GTesb5JX0s2d1oA")
	fmt.Println(err)
	// output:
	// false
}