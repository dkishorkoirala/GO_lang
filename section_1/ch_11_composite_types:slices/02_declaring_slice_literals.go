/*Declaring Slice Literals


Slice literals let you create and initialize slices in a single step. They're similar to array literals but without specifying the size.

Create a slice of strings using a slice literal:

colors := []string{"red", "blue", "green"}
fmt.Println(colors)
This displays:

[red blue green]
You can also create an empty slice:

emptySlice := []int{}
fmt.Println(emptySlice, len(emptySlice))
This displays:

[] 0

Challenge

In this challenge, you'll practice creating a slice literal in Go. Slice literals allow you to create and initialize a slice in a single statement.

Your task is to create a slice literal containing the names of four colors: "red", "blue", "green", and "yellow".*/

package main
import "fmt"

func main(){

	colors := [] string{"red", "blue", "green", "yellow"}

	fmt.Println("Colors:", colors)
	
	fmt.Println("Number of colors:", len(colors))
}