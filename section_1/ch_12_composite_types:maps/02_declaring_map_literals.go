/*Declaring Map Literals


Map literals let you create and initialize maps in a single statement. They're similar to slice literals but require both key and value for each entry.

Create a map literal with string keys and integer values:

ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Carol": 27,
}

fmt.Println(ages)
Result:

map[Alice:25 Bob:30 Carol:27]
You can also create an empty map literal:

emptyMap := map[string]int{}
fmt.Println(emptyMap)
Result:

map[]

Challenge

In this challenge, you'll practice creating a map literal in Go. Maps are collections of key-value pairs, similar to dictionaries in other languages.

Your task is to create a map literal that stores information about a few countries and their capitals. The map should have country names as keys (strings) and capital cities as values (also strings).*/

package main

import "fmt"

func main() {
	// TODO: Create a map literal called 'capitals' that maps countries to their capital cities
	// Include at least these three pairs:
	// "France" -> "Paris"
	// "Japan" -> "Tokyo"
	// "Brazil" -> "Brasilia"
	capitals := map[string]string{
		"France": "Paris",
		"Japan": "Tokyo",
		"Brazil": "Brasilia",
	}
	// This will print out the map
	fmt.Println(capitals)
	
	// This will print out the capital of Japan
	fmt.Println("The capital of Japan is:", capitals["Japan"])
}