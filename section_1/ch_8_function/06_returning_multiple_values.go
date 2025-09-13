/*Returning Multiple Values


Go functions can return multiple values at once, which is extremely useful for returning both results and error information.

Create a function that returns both a quotient and remainder:

func divide(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

func main() {
    q, r := divide(7, 3)
    fmt.Println("7 ÷ 3 =", q, "with remainder", r)
}
When you run this program, it outputs:

7 ÷ 3 = 2 with remainder 1
The function returns two values separated by commas, and we capture both values using multiple variables in the assignment.

Challenge

In this challenge, you'll practice returning multiple values from a function in Go.

We have a function called getPersonInfo that should return a person's name, age, and whether they are a student. Your task is to complete the function by adding the appropriate return statement.
*/

package main
import "fmt"

func getPersonInfo() (string, int, bool) {
	name := "Alex"
	age := 25
	isStudent := true
	
	return name, age, isStudent	
}

func main(){
	name, age, isStudent := getPersonInfo()

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Is Student: %t\n", isStudent)
}