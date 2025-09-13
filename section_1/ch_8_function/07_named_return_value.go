/*Named Return Values


In Go, you can name your return values in function declarations. This makes your code more readable and allows you to return by just using a plain return statement.

Create a function with named return values:

func divide(a, b int) (quotient int, remainder int) {
    quotient = a / b
    remainder = a % b
    return  // "naked" return
}

func main() {
    q, r := divide(7, 3)
    fmt.Println("7 ÷ 3 =", q, "with remainder", r)
}
The output is the same as before:

7 ÷ 3 = 2 with remainder 1
Named return values are initialized to their zero values and can be assigned directly in the function body.

Challenge

In this challenge, you'll practice using named return values in Go functions.

Complete the function getPersonInfo that returns a person's name, age, and whether they are a student. The function already has named return parameters, but you need to assign values to them.

The function should return the predefined values for a person named 'Alex' who is 25 years old and is a student.
*/
package main
import "fmt"

func getPersonInfo() (name string, age int, isStudent bool){
	name = "Alex"
	age = 25
	isStudent = true
	return
} 

func main() {
	name, age, isStudent := getPersonInfo()
	fmt.Printf("Name: %s\nAge: %d\nStudent: %t\n", name, age, isStudent)
}