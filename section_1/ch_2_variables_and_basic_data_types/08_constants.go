/*
Constants


Let's learn about constants in Go. Constants are values that cannot change during program execution.

First, let's declare some constants using the const keyword:

const pi = 3.14
const appName = "Go Calculator"
    
fmt.Println(pi)
fmt.Println(appName)
After executing this code, we'll see the values of our constants displayed on the screen:

3.14
Go Calculator
Unlike variables, constants cannot be changed after they're declared. If you try to assign a new value to a constant, your program will show an error. For example:

const pi = 3.14
pi = 5.6
// Error! because pi is const

Challenge

In this challenge, you'll practice using constants in Go. Constants are values that cannot be changed once they're defined.

Your task is to declare a constant named maxStudents with a value of 30, and then print it to the console.
*/

package main

import "fmt"

func main(){
	const maxStudents = 30
	fmt.Println("Maximum number of students allowed:", maxStudents)
}