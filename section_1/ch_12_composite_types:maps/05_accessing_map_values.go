/*Accessing Map Values


To retrieve values from a map, use the key in square brackets. Go also provides a way to check if a key exists.

Create a map of student scores:

students := map[string]int{
    "Emma": 92,
    "Jack": 78,
}

// Access a value using its key
emmaScore := students["Emma"]
fmt.Println("Emma's score:", emmaScore)
Result:

Emma's score: 92
Check if a key exists using the two-value assignment:

score, exists := students["Alex"]
fmt.Println("Score:", score, "Exists:", exists)
Result:

Score: 0 Exists: false

Challenge

In this challenge, you'll practice accessing values from a map in Go. A map of student grades has been created for you. Your task is to access and print the grade for a specific student.

The map contains student names as keys and their grades as values. You need to retrieve and print the grade for the student named 'Emma'.
*/

package main

import "fmt"

func main() {
	// A map of student grades
	studentGrades := map[string]string{
		"John":  "B+",
		"Emma":  "A-",
		"Sarah": "A",
		"Mike":  "C",
	}
	
	// TODO: Access and store Emma's grade in a variable called emmaGrade
	emmaGrade := studentGrades["Emma"]
	// TODO: Print Emma's grade with a message like: "Emma's grade is: A-"
	fmt.Println("Emma's grade is:", emmaGrade)
}