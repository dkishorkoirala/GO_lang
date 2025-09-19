/*Slicing Existing Slices/Arrays


You can create a new slice from an existing slice or array by specifying a range with the [start:end] syntax.

Create a slice of numbers:

numbers := []int{10, 20, 30, 40, 50}
fmt.Println("Original:", numbers)
Extract elements from index 1 to 3 (excluding 3):

slice1 := numbers[1:3]
fmt.Println("Slice [1:3]:", slice1)
Result:

Original: [10 20 30 40 50]
Slice [1:3]: [20 30]
You can also omit start or end index:

slice2 := numbers[:2]  // From start to index 2 (excluding 2)
slice3 := numbers[3:] // From index 3 to end
fmt.Println("Slice [:2]:", slice2)
fmt.Println("Slice [3:]:", slice3)
Result:

Slice [:2]: [10 20]
Slice [3:]: [40 50]

Challenge

In this challenge, you'll practice slicing an existing slice to create new sub-slices. You have a slice of fruits, and you need to create two new slices from it:

A slice containing only the first three fruits
A slice containing only the last two fruits
Then print both new slices to verify your work.*/

package main
import "fmt"

func main(){
	fruits := []string{"apple", "banana", "orange", "grape", "kiwi"}

	firstThree:= fruits[:3]
	lastTwo := fruits[3:]

	fmt.Println("First three fruits:", firstThree)
	fmt.Println("Last two fruits:", lastTwo)
}