/*Accessing Array Elements


To access array elements, use square brackets with an index number. Array indices start at 0.

Create an array of names:

names := [3]string{"Alex", "Beth", "Carlos"}
Access individual elements:

firstPerson := names[0]
secondPerson := names[1]
fmt.Println(firstPerson)
fmt.Println(secondPerson)
This displays:

Alex
Beth
You can also modify elements:

names[2] = "Charlie"
fmt.Println(names[2])
This outputs:

Charlie

Challenge

In this challenge, you'll practice accessing elements from an array in Go. We've created an array of fruits, and your task is to access specific elements and print them.

Complete the code to:
1. Access the first fruit in the array
2. Access the third fruit in the array
3. Print both fruits using fmt.Println*/

package main
import "fmt"
func main(){
    fruits := [5]string{"apple", "banana", "cherry", "date", "elderberry"}
	firstFruit := fruits[0]
	thirdFruit := fruits[2]
	fmt.Println("The first fruit is:", firstFruit)
    fmt.Println("The third fruit is:", thirdFruit)
}