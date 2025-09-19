/*Iterating Over Slices


To iterate over a slice in Go, you can use a for loop with the range keyword.

Create a slice of fruits:

fruits := []string{"Apple", "Banana", "Cherry"}
fmt.Println("Fruits:", fruits)
Iterate through the slice using range:

for index, value := range fruits {
    fmt.Printf("Index: %d, Value: %s\n", index, value)
}
Result:

Fruits: [Apple Banana Cherry]
Index: 0, Value: Apple
Index: 1, Value: Banana
Index: 2, Value: Cherry
If you only need the values, use the blank identifier (_):

for _, fruit := range fruits {
    fmt.Println(fruit)
}

Challenge

In this challenge, you'll practice iterating over a slice of fruits using the range keyword. Your task is to loop through the slice and print each fruit with its position in the list.

The slice is already defined for you. You need to complete the for loop using range to iterate through the slice and print each fruit with its position (starting from 1).
*/

package main
import "fmt"

func main(){
	fruits := []string{"Apple", "Banana", "Cherry", "Dragon fruit", "Elderberry"}

	for index, fruits := range(fruits){
		fmt.Printf("%d. %s\n",index+1, fruits)
	}
}