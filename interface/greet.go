package main

type Greet interface {
	GetName() string
	Hi() string
	Hello() string
}

type chinese struct {
	Name string
}

func (c *chinese) GetName() string {
	return c.Name
}

func (c *chinese) Hi() string {
	return"嗨!!"
}

func (c *chinese) Hello() string {
	return "哈囉!!"
}
