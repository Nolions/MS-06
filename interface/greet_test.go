package main

import (
	"github.com/stretchr/testify/mock"
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestSayHi(t *testing.T)  {
	m := new(chineseMock)

	assert.Equal(t, sayHi(m), "Hi!! Tom")
}

type chineseMock struct {
	mock.Mock
}

func (c *chineseMock) GetName() string{
	return "Tom"
}

func (c *chineseMock) Hi() string  {
	return "Hi!!"
}

func (c *chineseMock) Hello() string {
	return "Hello!!"
}