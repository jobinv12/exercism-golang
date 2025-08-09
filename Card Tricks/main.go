package main

import "fmt"

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	var sl []int

	return append(sl, 2, 6, 9)

}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {

	if index >= 0 && index <= len(slice) {
		return slice[index]
	}
	return -1

}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {

	if index >= 0 && index < len(slice) {
		slice[index] = value
		return slice
	}

	return append(slice, value)

}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {

	return append(values, slice...)

}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {

	if index >= 0 && index < len(slice) {
		return append(slice[:index], slice[index+1:]...)
	}
	return slice
}

func main() {

	var slice = [6]int{2, 4, 6, 8, 9, 10}
	index := 5
	value := 7

	fmt.Println(FavoriteCards())
	fmt.Println(GetItem(slice[:], index))
	fmt.Println(SetItem(slice[:], index, value))
	fmt.Println(PrependItems(slice[:], value))
	fmt.Println(RemoveItem(slice[:], index))

}
