/*Array Length with `len`


The len() function returns the length (number of elements) of an array.

Create an array of colors:

colors := [4]string{"Red", "Blue", "Green", "Yellow"}
Get the array length using len():

arrayLength := len(colors)
fmt.Println("The array has", arrayLength, "elements")
This displays:

The array has 4 elements

Challenge

In this challenge, you'll practice using the len() function to find the length of an array.

We've provided an array of favorite fruits. Your task is to find the length of this array using the len() function and store it in the numberOfFruits variable.*/

package main
import "fmt"

func main(){
	favoriteFruits := [5]string{"Apple", "Banana", "Orange", "Mango", "Strawberry"}
	numberOfFruits := len(favoriteFruits)

	fmt.Printf("There are %d fruits in the array\n", numberOfFruits)

}