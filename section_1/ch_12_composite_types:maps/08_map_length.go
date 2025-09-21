/*Map Length


The len function returns the number of key-value pairs in a map. This helps you determine how many entries a map contains.

Create a map of country populations:

populations := map[string]int{
    "USA": 331,
    "India": 1380,
    "China": 1441,
}

count := len(populations)
fmt.Println("Number of countries:", count)
The result shows the map's length:

Number of countries: 3
An empty map has length zero:

emptyMap := map[string]int{}
fmt.Println("Empty map length:", len(emptyMap))

Challenge

In this challenge, you'll practice using the len() function to find the number of key-value pairs in a map.

We've created a map called countries that contains some countries and their capitals. Your task is to use the len() function to find out how many countries are in the map, and then print that number.*/

package main

import "fmt"

func main() {
	// A map of countries and their capitals
	countries := map[string]string{
		"France":  "Paris",
		"Japan":   "Tokyo",
		"Brazil":  "Brasília",
		"Canada":  "Ottawa",
		"Egypt":   "Cairo",
	}
	
	// TODO: Use the len() function to get the number of countries in the map
	// and store it in a variable called countryCount
	countryCount := len(countries)
	
	// Print the number of countries
	fmt.Printf("The map contains %d countries\n", countryCount)
}