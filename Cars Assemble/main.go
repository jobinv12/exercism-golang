package main

import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {

	rate := (float64(productionRate) * successRate) / 100

	return float64(rate)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {

	rate := CalculateWorkingCarsPerHour(productionRate, successRate)
	produced := rate / 60

	return int(produced)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	tens := int(carsCount / 10)
	ones := int(carsCount % 10)

	return uint(tens*95000 + ones*10000)
}

func main() {

	productionRate := 1547
	successRate := 90
	carCount := 37

	fmt.Println(CalculateWorkingCarsPerHour(productionRate, float64(successRate)))
	fmt.Println(CalculateWorkingCarsPerMinute(productionRate, float64(successRate)))
	fmt.Println(CalculateCost(carCount))

}
