/*Recap - Creating Reusable Code


Challenge

You're building a simple calorie calculator for different activities.

Create a single function called calculateCalories that calculates how many calories are burned during an activity.

The function should:

Take an activity type (string), duration in minutes (int), and intensity level (float64)
Calculate and return the calories burned based on the activity type, duration, and intensity
Use these calorie rates per minute at normal intensity (1.0):

Running: 10 calories per minute
Swimming: 8 calories per minute
Cycling: 7 calories per minute
Any other activity: 5 calories per minute
The final calories should be calculated as: base calories × duration × intensity
*/
package main
import "fmt"

func calculateCalories(activity string, duration int, intensity float64) float64 {
	var baseCalories float64
	switch activity {
	case "running":
		baseCalories = 10.0
	case "swimming":
		baseCalories = 8.0
	case "cycling":
		baseCalories = 7.0
	default:
		baseCalories = 5.0
	}
	return baseCalories * float64(duration) * intensity
}

func main() {
    // Test the function with different activities
    fmt.Println("Running for 30 minutes at intensity 1.2:", calculateCalories("running", 30, 1.2))
    fmt.Println("Swimming for 45 minutes at intensity 1.0:", calculateCalories("swimming", 45, 1.0))
    fmt.Println("Cycling for 60 minutes at intensity 1.5:", calculateCalories("cycling", 60, 1.5))
    fmt.Println("Yoga for 60 minutes at intensity 0.8:", calculateCalories("yoga", 60, 0.8))
}