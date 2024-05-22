package main

import (
	"errors"
	"fmt"
	"os"
)

type Payer interface {
	Pay() (string, float64)
}

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual         Employee
	HourlyRate         float64
	HoursWorkedInAYear float64
	Review             map[string]interface{}
}

type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}

func (d Developer) FullName() string {
	return d.Individual.FirstName + " " + d.Individual.LastName
}

func (d Developer) Pay() (string, float64) {
	return d.FullName(), d.HourlyRate * d.HoursWorkedInAYear
}

func (m Manager) Pay() (string, float64) {
	fullName := m.Individual.FirstName + " " + m.Individual.LastName
	return fullName, m.Salary + (m.Salary * m.CommissionRate)
}

func payDetails(p Payer) {
	fullName, yearPay := p.Pay()
	fmt.Printf("Full Name: %s\nYearly Pay: %.2f\n", fullName, yearPay)
}

func OverallReview(i interface{}) (int, error) {
	switch v := i.(type) {
	case int:
		return v, nil
	case string:
		rating, err := ConvertReviewToInt(v)
		if err != nil {
			return 0, err
		}
		return rating, nil
	default:
		return 0, errors.New("Invalid Review")
	}
}

func ConvertReviewToInt(str string) (int, error) {
	switch str {
	case "Excellent":
		return 5, nil
	case "Good":
		return 4, nil
	case "Fair":
		return 3, nil
	case "Poor":
		return 2, nil
	case "Unsatisfactory":
		return 1, nil
	default:
		return 0, errors.New("Invalid Review")
	}
}

func (d Developer) ReviewRating() error {
	total := 0
	for _, v := range d.Review {
		rating, err := OverallReview(v)
		if err != nil {
			return err
		}
		total += rating
	}
	averageRating := float64(total) / float64(len(d.Review))
	fmt.Printf("%s got a review rating of %.2f\n", d.FullName(), averageRating)
	return nil
}

func main() {
	employeeReview := make(map[string]interface{})
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = "Good"
	employeeReview["Communication"] = "Excellent"
	employeeReview["ProblemSolving"] = "Fair"
	employeeReview["Dependability"] = "Poor"
	d := Developer{
		Individual: Employee{
			Id:        1,
			FirstName: "John",
			LastName:  "Doe",
		},
		HourlyRate:         35,
		HoursWorkedInAYear: 2080,
		Review:             employeeReview,
	}
	m := Manager{
		Individual: Employee{
			Id:        2,
			FirstName: "Jane",
			LastName:  "Doe",
		},
		Salary:         100000,
		CommissionRate: 0.07,
	}
	err := d.ReviewRating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	payDetails(d)
	payDetails(m)
}
