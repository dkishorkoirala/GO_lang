/*Adding and Updating Key-Value


To add or update values in a map, use the assignment operator with the key in square brackets.

Create a map of student scores:

students := make(map[string]int)

// Add new key-value pairs
students["Emma"] = 92
students["Jack"] = 78

fmt.Println(students)
Result:

map[Emma:92 Jack:78]
Update an existing value:

students["Jack"] = 85
fmt.Println(students)
Result:

map[Emma:92 Jack:85]

Challenge

In this challenge, you'll practice adding and updating values in a map. We have a map that stores the favorite fruits of different people. Your task is to add a new person's favorite fruit and update an existing person's favorite fruit.*/
package main

import (
    "fmt"
    "sort"
)

func main() {
	// Map of people and their favorite fruits
	favoriteFruits := map[string]string{
		"Alice": "Apple",
		"Bob":   "Banana",
		"Carol": "Cherry",
	}

	// TODO: Add a new entry for "David" with favorite fruit "Dragon Fruit"
	favoriteFruits["David"] = "Dragon Fruit"
	// TODO: Update Bob's favorite fruit to "Blueberry"
	favoriteFruits["Bob"] = "Blueberry"
	// Collect the names (keys) into a slice to sort them
	var names []string
	for person := range favoriteFruits {
		names = append(names, person)
	}
	sort.Strings(names)

	// Print the updated map in sorted order
	fmt.Println("Updated favorite fruits:")
	for _, person := range names {
		fruit := favoriteFruits[person]
		fmt.Printf("%s loves %s\n", person, fruit)
	}
}
