package main

import "fmt"

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0

	for i := 0; i < len(birdsPerDay); i++ {
		sum += birdsPerDay[i]
	}

	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {

	sum := 0

	for i := (week - 1) * 7; i < week*7; i++ {
		sum += birdsPerDay[i]
	}

	return sum

}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {

	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] += 1
	}

	return birdsPerDay

}

func main() {
	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	week := 2

	fmt.Println(TotalBirdCount(birdsPerDay))
	fmt.Println(BirdsInWeek(birdsPerDay, week))
	fmt.Println(FixBirdCountLog(birdsPerDay))
}
