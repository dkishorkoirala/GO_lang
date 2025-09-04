/*
Recap - Making Comparisons
Challenge

Complete the code to evaluate different comparison and logical expressions and print the results.
*/

package main
import "fmt"

func main(){
	age := 25
 	score := 85
  	isPremium := true
	
	isAdult := age >= 18
	isPassing := score >= 70
	isEligible := isAdult && isPremium
	needsImprovement := score < 70 || age < 20
	isRejected := !isPassing

	fmt.Println("Is Adult:", isAdult)
	fmt.Println("Is Passing:", isPassing)
	fmt.Println("Is Eligible:", isEligible)
	fmt.Println("Needs Improvement:", needsImprovement)
	fmt.Println("Is Rejected:", isRejected)
}