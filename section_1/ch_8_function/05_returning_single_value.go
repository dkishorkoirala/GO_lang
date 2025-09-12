/*Returning a Single Value


Functions can return values to the code that called them. This allows functions to compute results and give them back.

To create a function that returns a value, add the return type after the parameter list:

func multiply(x int, y int) int {
    result := x * y
    return result
}

func main() {
    product := multiply(4, 5)
    fmt.Println("4 × 5 =", product)
}
When you run this program, it outputs:

4 × 5 = 20
The return keyword sends the value back to where the function was called, and the returned value can be stored in a variable.

Challenge

In this challenge, you'll practice creating a function that returns a single value. You need to complete the getGreeting function that returns a greeting message based on the time of day.

The function should return "Good morning" if the hour is before 12, "Good afternoon" if the hour is between 12 and 17 (inclusive), and "Good evening" otherwise.
*/

package main
import "fmt"

func getGreeting(hour int) string {
	if (hour <12){
		return "Good morning"
	} else if (12<= hour && hour <=17){
		return "Good afternoon"
	}else{
		return "Good evening"
}}
func main(){
	morningHour := 8
    afternoonHour := 15
    eveningHour := 20
    
    fmt.Println("At", morningHour, "hours:", getGreeting(morningHour))
    fmt.Println("At", afternoonHour, "hours:", getGreeting(afternoonHour))
    fmt.Println("At", eveningHour, "hours:", getGreeting(eveningHour))
}