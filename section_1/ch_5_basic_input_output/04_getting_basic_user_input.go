/*
Getting Basic User Input


Let's learn how to get basic user input in Go. We'll create a simple program that asks for a name and age, then displays a greeting.

First, let's set up variables to store the user's name and age:

var name string
var age int
These variables will hold the information that the user types in. The string type is for text, and int is for whole numbers.

Next, let's ask the user for their name:

fmt.Print("Enter your name: ")
fmt.Scan(&name)
After executing this code, the text "Enter your name:" will show on the screen. The program will then wait for the user to type something and press Enter. Whatever they type will be stored in the name variable. The & symbol tells Go where to store the input.

Now, let's ask for the user's age:

fmt.Print("Enter your age: ")
fmt.Scan(&age)
Similar to the previous step, "Enter your age:" will show on the screen, and the program will wait for the user to type a number. This number will be stored in the age variable.

Finally, let's display a greeting using the information we collected:

fmt.Printf("Hello %s, you are %d years old!\n", name, age)
After executing this code, a personalized message will show on the screen.

For example, when you run this program and enter "Gopher" and "5", you'll see:

Enter your name: Gopher
Enter your age: 5
Hello Gopher, you are 5 years old!

Challenge

In this challenge, you'll practice getting basic user input in Go. You'll use the fmt.Scanln() function to read input from the user and store it in a variable.

Your task is to complete the program that asks the user for their name and then greets them.

Note: This challenge uses predefined variables instead of user input.
*/

package main
import "fmt"

func main(){
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s! Nice to meet you.\n", name)
}