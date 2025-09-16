/*Creating Slices with `make`


The make function creates slices with a specified length and capacity.

Create a slice of integers with length 3:

numbers := make([]int, 3)
fmt.Println(numbers)
This displays:

[0 0 0]
Create a slice with length 2 and capacity 5:

scores := make([]int, 2, 5)
fmt.Println(scores, len(scores), cap(scores))
This displays:

[0 0] 2 5

Challenge

In this challenge, you'll practice creating a slice using the make function. The make function allows you to create a slice with a specific length and capacity.

Your task is to create a slice of strings with a length of 3 and a capacity of 5 using make, then assign some values to it and print the slice.

After printing the slice, you should also print its length and capacity.*/

package main
import "fmt"

func main(){
	names := make([]string, 3, 5)
	names[0]="Alice"
	names[1]= "Bob"
	names[2]="Charlie"

	fmt.Println("Names:", names)

	fmt.Printf("Length: %d, Capacity: %d\n", len(names), cap(names))

}