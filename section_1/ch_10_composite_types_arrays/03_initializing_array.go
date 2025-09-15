/*Initializing Arrays


Initializing Arrays means setting initial values when creating an array.

Initialize an array with specific values using curly braces:

scores := [4]int{95, 87, 78, 92}
fmt.Println(scores)
This displays:

[95 87 78 92]
Use the ellipsis (...) to let Go count the elements for you:

colors := [...]string{"Red", "Green", "Blue"}
fmt.Println(colors)
fmt.Println("Array length:", len(colors))
This outputs:

[Red Green Blue]
Array length: 3

Challenge

In this challenge, you'll practice initializing arrays in Go. Your task is to create an array of 5 integers with the specific values provided, then print the array.

The code already has the necessary imports and a main function. You just need to initialize the array with the given numbers.*/
package main
import "fmt"

func main(){
	favouriteNumbers := [5]int{7,42,8,13,99}
	fmt.Println("My favorite numbers are:", favouriteNumbers)
}