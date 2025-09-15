/*Iterating Over Arrays


You can iterate through arrays using a for loop with an index.

Create an array of weekdays:

weekdays := [5]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
Iterate through the array using a for loop:

for i := 0; i < len(weekdays); i++ {
    fmt.Println(i, weekdays[i])
}
This displays each element with its index:

0 Monday
1 Tuesday
2 Wednesday
3 Thursday
4 Friday

Challenge

In this challenge, you'll practice iterating over an array of fruits using a for loop. The array is already defined for you. Your task is to complete the for loop to print each fruit in the array on a new line.*/

package main

import "fmt"

func main() {
	fruits := [5]string{"Apple", "Banana", "Orange", "Grape", "Mango"}
	
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])		
	}
}