/*Slice Length vs Capacity


In Go, length and capacity are two different properties of slices.

Create a slice with make:

numbers := make([]int, 3, 5)
fmt.Println("Length:", len(numbers))
fmt.Println("Capacity:", cap(numbers))
This displays:

Length: 3
Capacity: 5
Length is the number of elements currently in the slice. Capacity is the maximum number of elements the slice can hold before needing to grow.


Challenge

In this challenge, you'll explore the difference between a slice's len() and cap() functions. The len() function returns the current number of elements in a slice, while cap() returns the capacity (maximum number of elements before reallocation is needed).

Your task is to complete the code to print both the length and capacity of the given slice.*/

package main
import "fmt"

func main(){
	numbers := make([]int, 3, 5)

	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))
}