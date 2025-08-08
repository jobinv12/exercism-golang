package main

import "fmt"

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {

	if name != "" {
		return "One for " + name + ", one for me."
	}
	return "One for you, one for me."
}
func main() {

	name := "gopher"

	fmt.Println(ShareWith(name))
}
