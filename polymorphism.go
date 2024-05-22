package main

import "fmt"

type Speaker interface {
	Speak() string
}

type cat struct{}
type dog struct{}
type person struct {
	name string
}

func (c cat) Speak() string {
	return "Purr Meow"
}

func (d dog) Speak() string {
	return "Woof Woof"
}

func (p person) Speak() string {
	return "Hello, my name is " + p.name
}

func catSpeack(c cat) {
	fmt.Println(c.Speak())
}

func dogSpeack(d dog) {
	fmt.Println(d.Speak())
}

func personSpeak(p person) {
	fmt.Println(p.Speak())
}

func main() {
	c := cat{}
	d := dog{}
	p := person{name: "John"}
	catSpeack(c)
	dogSpeack(d)
	personSpeak(p)
}
