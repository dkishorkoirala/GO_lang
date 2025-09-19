/*Copying Slices


The copy function in Go lets you copy elements from one slice to another.

Create a source slice and a destination slice:

source := []int{1, 2, 3}
dest := make([]int, 2)  // Destination with length 2

// Copy from source to destination
copied := copy(dest, source)

fmt.Println("Source:", source)
fmt.Println("Destination:", dest)
fmt.Println("Number of elements copied:", copied)
Result:

Source: [1 2 3]
Destination: [1 2]
Number of elements copied: 2
The copy function copies only as many elements as will fit in the destination slice.

Challenge

In this challenge, you'll practice using the copy function to copy elements from one slice to another.

You have two slices: source containing fruit names and destination which is an empty slice with capacity for 3 elements. Your task is to copy the elements from source to destination using the copy function.

After copying, print the contents of the destination slice.*/

package main
import "fmt"

func main(){
	source := []string{"apple", "banana", "cherry", "date"}
	destination := make([]string,3)

	copy(destination, source)
	fmt.Println("Destination slice:", destination)
}