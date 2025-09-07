/*Recap - Making Decisions


Challenge

Complete the code to implement different conditional statements and print the results.
*/

package main 
import "fmt"

func main(){
	score := 85
  	grade := ""

	if score >= 90{
		grade = "A"
	}else if score >=80{
		grade = "B"
	}else if score >=70{
		grade ="C"
	}else{
		grade = "F"
	}
	fmt.Println("Grade:",grade)
	switch grade{
	case "A":
		fmt.Println("Excellent work!")
	case "B":
		fmt.Println("Good job!")
	case "C":
		fmt.Println("Satisfactory.")
	default:
		fmt.Println("You need to improve.")
	}

}