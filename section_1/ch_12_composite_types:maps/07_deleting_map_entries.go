/*Deleting Map Entries


To delete an entry from a map in Go, use the built-in delete function. This removes a key-value pair from the map.

Create a map of scores:

scores := map[string]int{
    "Alice": 95,
    "Bob": 82,
    "Charlie": 78,
}
fmt.Println(scores)
Delete Bob's score from the map:

delete(scores, "Bob")
fmt.Println(scores)
Result:

map[Alice:95 Charlie:78]
Deleting a key that doesn't exist is safe and has no effect:

delete(scores, "David") // No error

Challenge

In this challenge, you'll practice deleting entries from a map in Go.

We have a map called inventory that contains items and their quantities. Your task is to delete the item 'pencil' from the inventory map using the delete function.

After deleting the entry, the program will check if the item still exists in the map and print the results.*/

package main

import "fmt"

func main() {
	// Map of inventory items and their quantities
	inventory := map[string]int{
		"pen":    15,
		"pencil": 10,
		"paper":  25,
		"eraser": 5,
	}

	fmt.Println("Initial inventory:", inventory)

	// TODO: Delete the 'pencil' entry from the inventory map
	delete(inventory, "pencil")

	fmt.Println("Updated inventory:", inventory)

	// Check if 'pencil' still exists in the map
	_, exists := inventory["pencil"]
	fmt.Println("Does 'pencil' exist in inventory?", exists)
}
