// +build !aws

package main

import (
	"fmt"
	"os"
	"strconv"
	. "github.com/paul-nelson-baker/ball-clock-simulator/support"
)

func main() {
	fmt.Println("Instructions: Enter a single number for mode 1, or two numbers separated by a space for mode 2")
	fmt.Println("              Ball count must be in the range [27, 127] inclusive.")
	fmt.Println("Mode 1: Ball count only")
	fmt.Println("Mode 2: Ball count and iteration count")
	fmt.Println("Use CTRL+C to exit")
	for {
		fmt.Print("Enter your input here: ")
		if userInputValues, err := getValidUserInputFromStdIn(); err != nil {
			// Anything goes wrong, print out the error and re-ask for input
			fmt.Println(err.Error())
		} else {
			switch len(userInputValues) {
			case 1:
				//If mode 1, calculate how many days before we reset
				ballCount := userInputValues[0]
				fmt.Println("Mode 1::Working with ballcount: " + strconv.Itoa(ballCount))
				fmt.Print(CalculateDaysUntilResetString(ballCount))
			case 2:
				// if mode 2, perform the number of iterations necessary
				// and then print out the state of the clock
				ballCount := userInputValues[0]
				iterations := userInputValues[1]
				fmt.Println("Mode 2::Working with ballcount: " + strconv.Itoa(ballCount))
				fmt.Println("Mode 2::Working with iterations: " + strconv.Itoa(iterations))
				ballClock := NewBallClock(ballCount)
				ballClock.TickMinutes(iterations)
				fmt.Println(ballClock.JsonString())
			}
		}
		fmt.Println()
	}
}

func getValidUserInputFromStdIn() ([]int, error) {
	return GetValidUserInput(os.Stdin)
}
