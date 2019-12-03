package jwt

import "fmt"

func Example() {
	jwtToken := signing(&user{
		id:      2,
		account: "sam78",
	})
	token, _ := validating(jwtToken)
	fmt.Println(token["id"])
	fmt.Println(token["account"])

	// output:
	// 2
	// sam78
}

func ExampleValidating_Fail_Token_Expired() {
	_, err := validating("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50Ijoic2FtNzgiLCJleHAiOjE1NzUyNjI4MjAsImlkIjoyfQ.F2rfrP8rtenUnlsaSFX-AiVUq1h-2bVnXFFWFvYu-fM")
	if err.Code != 0 {
		fmt.Println(err.Code)
		fmt.Println(err.Message)
	}

	// output:
	// 101
	// Token Expired
}

func ExampleValidating_Fail_Signature_Invalid() {
	_, err := validating("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.ViFKe4FNOkwN0DYvlfhEqkPUms4DbqrE7L2zhIT8tnU")
	if err.Code != 0 {
		fmt.Println(err.Code)
		fmt.Println(err.Message)
	}

	// output:
	// 102
	// Token Invalid
}

func ExampleValidating_Fail() {
	_, err := validating("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")
	if err.Code != 0 {
		fmt.Println(err.Code)
		fmt.Println(err.Message)
	}
	// output:
	// 103
	// Token Format error
}
