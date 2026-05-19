package main

import (
	"fmt"
)

func showMenu() {
	fmt.Println("\n===== Result Checker =====")
	fmt.Println("1. Grade Check")
	fmt.Println("2. Pass/Fail Check")
	fmt.Println("3. Exit")
	fmt.Print("Enter your choice: ")
}

func validateMarks(marks int) error {
	if marks < 0 {
		return fmt.Errorf("marks cannot be less than 0")
	}

	if marks > 100 {
		return fmt.Errorf("marks cannot be greater than 100")
	}

	return nil
}

func getMarksInput() int {
	var marks int

	for {
		fmt.Print("Enter marks (0-100): ")
		fmt.Scan(&marks)

		err := validateMarks(marks)

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		return marks
	}
}

func checkGrade(marks int) {
	if marks >= 80 {
		fmt.Println("You've got GPA 5.00")
	} else if marks >= 60 {
		fmt.Println("You've got GPA 4.00")
	} else if marks >= 33 {
		fmt.Println("You've got GPA 3.00")
	} else {
		fmt.Println("You've got F")
	}
}

func checkPassFail(marks int) {
	if marks >= 33 {
		fmt.Println("Congratulations! You've passed.")
	} else {
		fmt.Println("Sorry! You've failed.")
	}
}

func main() {
	running := true

	for running {
		showMenu()

		var choice int
		fmt.Scan(&choice)

		switch choice {

		case 1:
			marks := getMarksInput()
			checkGrade(marks)

		case 2:
			marks := getMarksInput()
			checkPassFail(marks)

		case 3:
			fmt.Println("Exiting program...")
			running = false

		default:
			fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
		}
	}
}
