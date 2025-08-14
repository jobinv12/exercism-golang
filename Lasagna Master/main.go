package main

import "fmt"

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTime int) int {

	if avgTime == 0 {
		avgTime = 2
	}
	return len(layers) * avgTime

}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {

	for _, l := range layers {
		switch l {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}

	return noodles, sauce

}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) []string {

	myList[len(myList)-1] = friendsList[len(friendsList)-1]

	return myList

}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {

	toCook := make([]float64, len(quantities))

	for i, quantity := range quantities {
		toCook[i] = quantity * float64(portions) / 2
	}

	return toCook

}

func main() {
	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	avgTime := 3
	myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}
	quantities := []float64{1.2, 3.6, 10.5}
	portion := 4

	fmt.Println(PreparationTime(friendsList, avgTime))
	fmt.Println(Quantities(friendsList))
	fmt.Println(AddSecretIngredient(friendsList, myList))
	fmt.Println(ScaleRecipe(quantities, portion))
}
