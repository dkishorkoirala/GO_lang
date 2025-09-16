/*Accessing Slice Elements


Accessing slice elements works just like arrays, using square brackets with an index position (starting from 0).

Create a slice of fruits:

fruits := []string{"Apple", "Banana", "Cherry"}
fmt.Println(fruits[0])  // First element
fmt.Println(fruits[1])  // Second element
fmt.Println(fruits[2])  // Third element
This displays:

Apple
Banana
Cherry
You can also modify elements by assigning new values:

fruits[1] = "Blueberry"
fmt.Println(fruits)
Result:

[Apple Blueberry Cherry]

Challenge

In this challenge, you'll practice accessing elements from a slice in Go.

We've created a slice of fruits for you. Your task is to access specific elements from this slice and print them.*/

package main
import "fmt"

func main(){
	fruits := []string{"apple", "banana", "cherry", "date", "elderberry"}
	firstFruit := fruits[0]
	thirdFruit := fruits[2]
	fmt.Println("The first fruit is:", firstFruit)
	fmt.Println("The third fruit is:", thirdFruit)
}