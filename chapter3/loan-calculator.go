package main

import (
	"errors"
	"fmt"
)

const (
	goodScore      = 450
	lowScoreRatio  = 10
	goodScoreRatio = 20
)

var (
	ErrCreditScore = errors.New("Invalid credit score")
	ErrIncome      = errors.New("Income invalid")
	ErrLoanAmount  = errors.New("Loan amount invalid")
	ErrLoanTerm    = errors.New("Loan term invalid")
)

func checkLoan(creditScore int, income float64, loanAmount float64, loanTerm float64) error {
	interest := 20.0

	if creditScore >= goodScore {
		interest = 15.0
	}

	if creditScore < 1 {
		return ErrCreditScore
	}

	if income < 1 {
		return ErrIncome
	}

	if loanAmount < 1 {
		return ErrLoanAmount
	}

	if loanTerm < 1 || int(loanTerm)%12 != 0 {
		return ErrLoanTerm
	}

	rate := interest / 100
	payment := ((loanAmount * rate) / loanTerm) + (loanAmount / loanTerm)

	totalInterest := (payment * loanTerm) - loanAmount
	approved := false

	if income > payment {
		ratio := (payment / income) * 100
		if creditScore >= goodScore && ratio < goodScoreRatio {
			approved = true
		} else if ratio < lowScoreRatio {
			approved = false
		}
	}

	fmt.Println("Credit Score    :", creditScore)
	fmt.Println("Income          :", income)
	fmt.Println("Loan Amount     :", loanAmount)
	fmt.Println("Loan Term       :", loanTerm)
	fmt.Println("Monthly Payment :", payment)
	fmt.Println("Rate            :", interest)
	fmt.Println("Total Cost      :", totalInterest)
	fmt.Println("Approved        :", approved)
	fmt.Println("")

	return nil
}

func main() {
	// Approved
	fmt.Println("Applicant 1")
	fmt.Println("-----------")
	err := checkLoan(500, 1000, 1000, 24)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Denied
	fmt.Println("Applicant 2")
	fmt.Println("-----------")
	err = checkLoan(350, 1000, 10000, 12)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
