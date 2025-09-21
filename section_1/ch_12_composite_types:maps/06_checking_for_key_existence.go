/*Checking for Key Existence


In Go, you can check if a key exists in a map using the two-value assignment form. This is important because accessing a non-existent key returns the zero value, which might be misleading.

Create a map of country codes:

countries := map[string]string{
    "US": "United States",
    "CA": "Canada",
}

// Check if "UK" exists in the map
value, exists := countries["UK"]
fmt.Println("Value:", value, "Exists:", exists)
Result:

Value:  Exists: false
The exists boolean tells you whether the key was found. Use this in conditionals:

if _, exists := countries["US"]; exists {
    fmt.Println("US is in the map")
}
Result:

US is in the map

Challenge

In this challenge, you'll practice checking if a key exists in a map. You have a map of student grades, and you need to check if a specific student's grade exists in the map.

Complete the code to check if the student 'Emma' exists in the grades map, and print an appropriate message.*/

package main

import "fmt"

func main() {
	// Map of student grades
	grades := map[string]int{
		"John":  85,
		"Sarah": 92,
		"Mike":  78,
		"Lisa":  95,
	}

	// Student to check
	studentToCheck := "Emma"

	// TODO: Check if studentToCheck exists in the grades map
	_, exists := grades[studentToCheck]

	// and store the result in a variable called 'exists'
	// Hint: Use the comma-ok idiom


	// Print the result
	if exists {
		fmt.Printf("%s's grade exists in the map\n", studentToCheck)
	} else {
		fmt.Printf("%s's grade does not exist in the map\n", studentToCheck)
	}
}