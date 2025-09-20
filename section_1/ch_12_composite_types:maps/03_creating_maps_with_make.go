/*Creating Maps with `make`


The make function provides another way to create maps, especially when you don't need to initialize values immediately.

Create an empty map with make:

scores := make(map[string]int)
fmt.Println(scores)
Result:

map[]
Add elements to the map after creation:

scores["Alice"] = 95
scores["Bob"] = 80
fmt.Println(scores)
Result:

map[Alice:95 Bob:80]

Challenge

In this challenge, you'll practice creating a map using the make function in Go. You'll create a map that stores the favorite colors of different people.

Your task is to create a map using make and then add some key-value pairs to it.*/

package main

import "fmt"

func main() {
	// TODO: Create a map called 'favoriteColors' using make
	// The keys should be strings (names) and the values should also be strings (colors)
	favoriteColors := make(map[string]string)

	// Adding some key-value pairs to the map
	favoriteColors["Alice"] = "Blue"
	favoriteColors["Bob"] = "Green"
	favoriteColors["Charlie"] = "Red"
	
	// Print the map
	fmt.Println("Favorite colors:", favoriteColors)
	
	// Print Bob's favorite color
	fmt.Println("Bob's favorite color is", favoriteColors["Bob"])
}