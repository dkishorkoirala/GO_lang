/*
Logical AND Operator


The logical AND operator (&&) returns true only when both conditions are true.

Here is a program that checks if a number is both positive and even:

num := 6
isPositive := num > 0
isEven := num % 2 == 0
result := isPositive && isEven
fmt.Println(result)
This outputs:

true
You can also use the operator directly in conditions:

age := 25
hasLicense := true
canDrive := age >= 18 && hasLicense
fmt.Println(canDrive)
This outputs:

true

Challenge

In this challenge, you'll practice using the logical AND operator (&&) in Go.

You have variables representing whether a person meets different eligibility criteria for a job. A person is qualified for the job if they have both the required experience AND the necessary education.

Your task is to use the logical AND operator to check if the candidate is qualified, and store the result in the isQualified variable.
*/

package main
import "fmt"

func main(){
	hasExperience := true
	hasEducation := true

	isQualified := hasExperience && hasEducation
	fmt.Println("Does the candidate qualify for the job?", isQualified)
}