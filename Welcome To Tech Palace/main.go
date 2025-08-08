package main

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {

	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {

	return strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {

	newMsg := strings.Replace(oldMsg, "*", "", -1)

	return strings.TrimSpace(newMsg)

}

func main() {

	customer := "Gopher"
	welcomeMsg := "Welcome"
	numStars := 10
	oldMsg := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`
	fmt.Println(WelcomeMessage(customer))
	fmt.Println(AddBorder(welcomeMsg, numStars))
	fmt.Println(CleanupMessage(oldMsg))

}
