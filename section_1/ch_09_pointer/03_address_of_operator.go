/*The Address-Of Operator


The address-of operator (&) gets the memory address of a variable. It's how we create pointers in Go.

Create a variable and get its memory address:

func main() {
    number := 42
    addressOfNumber := &number
    
    fmt.Println("Value:", number)
    fmt.Println("Address:", addressOfNumber)
}
The output shows the value and its memory address:

Value: 42
Address: 0xc000018030
The address (like 0xc000018030) is a hexadecimal number representing the memory location where the value is stored. This address will be different each time you run the program.

Challenge

In this challenge, you'll practice using the address-of operator (&) in Go. The address-of operator returns the memory address of a variable, which is where the variable's value is stored in the computer's memory.

Your task is to use the address-of operator to get the memory address of the provided variable and store it in the pointer variable that's already declared.
*/

package main
import "fmt"

func main(){
	name := "Gopher"

	var namePointer *string

	namePointer = &name
	fmt.Printf("The value at that memory address is: %v\n", *namePointer)

}