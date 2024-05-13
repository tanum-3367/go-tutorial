package main

import (
	"fmt"
	"unicode"
)

func passwordChecker(pw string) bool {
	pwR := []rune(pw)

	if len(pwR) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, r := range pwR {
		if unicode.IsUpper(r) {
			hasUpper = true
		}
		if unicode.IsLower(r) {
			hasLower = true
		}
		if unicode.IsNumber(r) {
			hasNumber = true
		}
		if unicode.IsPunct(r) || unicode.IsSymbol(r) {
			hasSymbol = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {
	if passwordChecker("This!I5A") {
		fmt.Println("password good")
	} else {
		fmt.Println("password bad")
	}
}
