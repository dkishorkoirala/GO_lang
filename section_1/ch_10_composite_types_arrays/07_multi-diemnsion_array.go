/*Multi-dimensional Arrays


Multi-dimensional arrays are arrays that contain other arrays as elements. They're useful for representing grids or tables.

Create a 2×3 array (2 rows, 3 columns) to represent a grid of numbers:

grid := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}
Access elements using two indices - first for row, then for column:

element := grid[0][2]
fmt.Println(element)
This displays the element at row 0, column 2:

3

Challenge

In this challenge, you'll work with a multi-dimensional array that represents a small grid of numbers. Your task is to access and print a specific element from this 2D array.

The grid represents a simple 3x3 game board with numbers:

1 2 3
4 5 6
7 8 9
Your job is to print the value at row 1, column 2 (which should be the number 6).*/

package main
import "fmt"

func main(){
	grid := [3][3]int{
		{1,2,3},{4,5,6},{7,8,9},
	}
	
	fmt.Println(grid[1][2])
}