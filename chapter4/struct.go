package main

import "fmt"

type user struct {
	name    string
	age     int
	balance float64
	member  bool
}

func getUsers() []user {
	u1 := user{
		name:    "John",
		age:     25,
		balance: 100.0,
		member:  true,
	}
	u2 := user{
		age:  19,
		name: "Hanvy",
	}
	u3 := user{
		"Bob",
		30,
		200.0,
		false,
	}

	var u4 user
	u4.name = "Alice"
	u4.age = 22
	u4.balance = 150.0
	u4.member = true

	return []user{u1, u2, u3, u4}
}

func main() {
	users := getUsers()
	for i := 0; i < len(users); i++ {
		fmt.Printf("User %v: %#v\n", i, users[i])
	}
}
