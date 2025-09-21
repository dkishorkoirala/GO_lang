/*Iterating Over Maps


The range keyword lets you iterate over all key-value pairs in a map. It returns both the key and value for each entry.

Create a map of fruit inventory:

inventory := map[string]int{
    "apples": 5,
    "bananas": 8,
    "oranges": 12,
}

for fruit, quantity := range inventory {
    fmt.Printf("%s: %d\n", fruit, quantity)
}
The result shows each key-value pair:

apples: 5
bananas: 8
oranges: 12
If you only need the keys, you can omit the second variable:

for fruit := range inventory {
    fmt.Println(fruit)
}

Challenge
In this challenge, you'll practice iterating over a map using the range keyword in Go. A map of student grades has been created for you. Your task is to iterate through the map and print each student's name and their grade.

Complete the code to iterate through the studentGrades map and print each student's name and grade in the format: Student: [name], Grade: [grade]*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	// A map of student grades
	studentGrades := map[string]string{
		"Alice":  "A",
		"Bob":    "B",
		"Charlie": "B+",
		"David":  "A-",
		"Eva":    "C",
	}

	// Collect the student names (keys) into a slice to sort them
	var names []string
	for name := range studentGrades {
		names = append(names, name)
	}
	sort.Strings(names)


	// TODO: Iterate through the sorted list of student names
	// and print each student's name and grade in the format:
	// "Student: [name], Grade: [grade]"
	for _, name := range names {
		grade := studentGrades[name]
		fmt.Printf("Student: %s, Grade: %s\n", name, grade)
	}
}
