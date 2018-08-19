package support

import (
	"io"
	"bufio"
	"strings"
	"strconv"
	"errors"
)

var (
	// Reusable errors that we can specifically check for later, if necessary, by reference
	InvalidInputString = errors.New("invalid input. input must be either one number, or a two numbers separated by a single space")
	InvalidInputCount  = errors.New("invalid number of inputs. cannot be less than one item or more than two")
	InvalidBallCount   = errors.New("invalid ball count. cannot be less than 27 or more than 127")
)

func GetValidUserInput(in io.Reader) ([]int, error) {
	// Get a single line from the user and split them by space
	reader := bufio.NewReader(in)
	userInputLine, _ := reader.ReadString('\n')
	userInputItems := strings.Fields(userInputLine)

	// Iterate over the items returned from the user and turn
	// them to integers (regardless of how many)
	var results []int
	for _, value := range userInputItems {
		// If a value isn't a valid numerical string, we short-circuit
		// and return an empty slice along with an error we can print
		if number, err := strconv.Atoi(strings.TrimSpace(value)); err != nil {
			return []int{}, InvalidInputString
		} else {
			results = append(results, number)
		}
	}
	// Now we validate we have the correct number of items.
	// We only have mode 1 and 2, so we enforce that via length
	if len(results) < 1 || len(results) > 2 {
		return []int{}, InvalidInputCount
	}
	// Now that we have validated the user-date is structurally
	// correct we can enforce that it is also logically correct.
	// Second ball must be in the [27, 127] range. I believe inclusively.
	if results[0] < 27 || results[0] > 127 {
		return []int{}, InvalidBallCount
	}
	// We've gone through the gamut of various things that can go
	// wrong, but we've proven the data is sanitary.
	return results, nil
}
