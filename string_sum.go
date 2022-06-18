package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func sum(a, b string) (string, error) {
	first, err1 := strconv.Atoi(a)
	if err1 != nil {
		return "", fmt.Errorf("not valid input: %w", err1)
	}

	second, err2 := strconv.Atoi(b)
	if err2 != nil {
		return "", fmt.Errorf("not valid input: %w", err2)
	}
	return strconv.Itoa(first + second), nil
}

func StringSum(input string) (output string, err error) {
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, "\t", "")
	input = strings.ReplaceAll(input, "\n", "")

	if input == "" {
		return "", fmt.Errorf("empty string provided: %w", errorEmptyInput)
	}
	operands := strings.Split(input, "+")

	if len(operands) == 1 {
		if strings.Count(input, "-") > 2 || strings.Count(input, "-") == 0 {
			return "", fmt.Errorf("unexpected number of operands: %w", errorNotTwoOperands)
		}
		lastIndex := strings.LastIndex(input, "-")
		if len(input)-1 == lastIndex {
			return "", fmt.Errorf("unexpected number of operands: %w", errorNotTwoOperands)
		}
		return sum(input[:lastIndex], input[lastIndex:])
	}

	if len(operands) != 2 {
		return "", fmt.Errorf("two many operands: %w", errorNotTwoOperands)
	}

	return sum(operands[0], operands[1])
}
