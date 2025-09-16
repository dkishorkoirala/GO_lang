/*Introduction to Slices


Slices are flexible, dynamic arrays in Go. Unlike arrays with fixed size, slices can grow or shrink.

Create a slice using square brackets without a size, followed by the type:

numbers := []int{1, 2, 3, 4, 5}
fmt.Println(numbers)
This displays the slice contents:

[1 2 3 4 5]
Access slice elements using index notation (starting from 0):

firstNumber := numbers[0]
fmt.Println(firstNumber)
This displays the first element:

1*/