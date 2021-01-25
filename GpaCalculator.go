package main

import (
	"fmt"
)

// returns the average of the gpa from each class
func getAverage(nums []float64) float64 {
	average := 0.0
	for _, val := range nums {
		average += val
	}
	average /= float64(len(nums))
	return average

}

// converts to gpa form (4.0, 3.0, etc)
func gradesToGradingScale(nums []int) []float64 {
	var gpaSlice []float64
	for _, val := range nums {
		newGPA := 4.0
		if val < 90 {
			newGPA = 3.0
		}
		if val < 80 {
			newGPA = 2.0
		}
		if val < 70 {
			newGPA = 1.0
		}
		if val < 65 {
			newGPA = 0.0
		}
		gpaSlice = append(gpaSlice, newGPA)
	}
	return gpaSlice
}

func main() {
	fmt.Println("GPA Calculator 1.0")

	var numClasses int

	fmt.Println("Type the number of classes you are taking.")
	fmt.Scanln(&numClasses)

	grades := make([]int, numClasses)

	for i := 0; i < numClasses; i++ {
		fmt.Printf("Type the grade (0-100) for class #%d.\n", i+1)
		fmt.Scanln(&grades[i])
	}
	fmt.Println(getAverage(gradesToGradingScale(grades)))

}
