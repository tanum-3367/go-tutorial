package main

import "fmt"

type Speaker interface {
	Speak() string
}

type cat struct {
}

func (c cat) Speak() string {
	return "Purr Meow"
}

func (c cat) Greeting() string {
	return "hello"
}

func main() {
	c := cat{}
	fmt.Println(c.Speak())
	fmt.Println(c.Greeting())
}
