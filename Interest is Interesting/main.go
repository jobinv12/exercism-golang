package main

import "fmt"

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {

	switch {
	case balance < 0:
		return 3.213
	case balance < 1000:
		return 0.5
	case balance < 5000:
		return 1.621
	default:
		return 2.475
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {

	interest := balance * float64(InterestRate(balance)) / 100

	return float64(interest)

}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {

	interest := Interest(balance)

	return balance + interest

}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {

	years := 0
	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		years++
	}
	return years

}

func main() {

	balance := 2000
	targetBalance := 2500

	fmt.Println(InterestRate(float64(balance)))
	fmt.Println(Interest(float64(balance)))
	fmt.Println(AnnualBalanceUpdate(float64(balance)))
	fmt.Println(YearsBeforeDesiredBalance(float64(balance), float64(targetBalance)))
}
