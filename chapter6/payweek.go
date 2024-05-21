package main

import (
	"errors"
	"fmt"
	"strings"
)

type DirectDeposit struct {
	lastName      string
	firstName     string
	bankName      string
	routingNumber int
	accountNumber int
}

var ErrInvalidLastName = errors.New("Invalid last name")
var ErrInvalidRoutingNum = errors.New("Invalid routing number")

func (dd *DirectDeposit) validateRoutingNumber() error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from: ", r)
		}
	}()

	if dd.routingNumber < 100 {
		panic(ErrInvalidRoutingNum)
	}
	return nil
}

func (dd *DirectDeposit) validateLastName() error {
	dd.lastName = strings.TrimSpace(dd.lastName)
	if len(dd.lastName) == 0 {
		return ErrInvalidLastName
	}
	return nil
}

func (dd *DirectDeposit) report() {
	fmt.Println(strings.Repeat("*", 80))
	fmt.Println("Last Name: ", dd.lastName)
	fmt.Println("First Name: ", dd.firstName)
	fmt.Println("Bank Name: ", dd.bankName)
	fmt.Println("Routing Number: ", dd.routingNumber)
	fmt.Println("Account Number: ", dd.accountNumber)
}

func main() {
	dd := DirectDeposit{

		lastName:      "  ",
		firstName:     "Abe",
		bankName:      "WilkesBooth Inc",
		routingNumber: 17,
		accountNumber: 1809,
	}

	err := dd.validateRoutingNumber()
	if err != nil {
		fmt.Println(err)
	}
	err = dd.validateLastName()
	if err != nil {
		fmt.Println(err)
	}

	dd.report()
}
