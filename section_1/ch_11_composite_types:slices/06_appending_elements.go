/*Appending Elements


The append function adds elements to the end of a slice, automatically handling capacity changes.

Create a slice of numbers:

numbers := []int{1, 2, 3}
fmt.Println(numbers)
Append a new element:

numbers = append(numbers, 4)
fmt.Println(numbers)
Result:

[1 2 3]
[1 2 3 4]
You can append multiple elements at once:

numbers = append(numbers, 5, 6, 7)
fmt.Println(numbers)
Result:

[1 2 3 4 5 6 7]

Challenge

In this challenge, you'll practice appending elements to a slice in Go. We've created a slice of fruits, and your task is to append two more fruits to it using the append() function.*/


package main
import "fmt"

func main(){
	fruits := []string{"apple", "banana", "orange"}

	fruits = append(fruits, "grape", "kiwi")
	fmt.Println("My fruit collection:", fruits)
}
