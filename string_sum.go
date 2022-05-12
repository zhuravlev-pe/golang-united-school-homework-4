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

func StringSum(input string) (output string, err error) {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return "", errorEmptyInput
	}
	operands := strings.Split(input, "+")
	if len(operands) != 2 {
		return "", errorNotTwoOperands
	}

	op1, err := parseOperand(operands[0])
	if err != nil {
		return "", err
	}

	op2, err := parseOperand(operands[1])
	if err != nil {
		return "", err
	}

	result := op1 + op2

	return strconv.Itoa(result), nil
}

func parseOperand(operand string) (int, error) {
	operand = strings.TrimSpace(operand)
	result, err := strconv.Atoi(operand)
	if err != nil {
		return 0, fmt.Errorf("bad operand %q: %w", operand, err)
	}
	return result, nil
}
