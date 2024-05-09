package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var helloList = []string{
	"Hello, world",
	"Καλημέρα κόσμε",
	"こんにちは世界",
	"سلام دنیا",
	"Привет, мир",
}

func hello(index int) (string, error) {
	if index < 0 || index > len(helloList)-1 {
		return "", errors.New("out of range: " + strconv.Itoa(index))
	}

	return helloList[index], nil
}

func main() {
	// seed random number generator using current time
	rand.NewSource(time.Now().UnixNano())
	// generate a random number in the range of our list
	index := rand.Intn(len(helloList))
	// call a function and receive multiple return values
	msg, err := hello(index)
	// handle any errors
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}
