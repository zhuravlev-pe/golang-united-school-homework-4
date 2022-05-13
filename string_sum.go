package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

var noMoreOperands = errors.New("no more operands")

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
	//fmt.Printf("input: %q\n", input)
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return "", fmt.Errorf("empty input: %w", errorEmptyInput)
	}
	var operands []operand
	runes := []rune(input)
	pos := 0
	for {
		operand, err := getOperand(runes, pos)
		if err != nil {
			if err == noMoreOperands {
				break
			}
			return "", err
		}
		if len(operands) > 0 {
			if !operand.hasSign {
				return "",
					fmt.Errorf("missing sign between operands %d and %d",
						len(operands), len(operands)+1)
			}
		}
		operands = append(operands, operand)
		pos = operand.posEnd
	}

	if len(operands) != 2 {
		return "", fmt.Errorf("bad input %q: %w", input, errorNotTwoOperands)
	}

	result := operands[0].value + operands[1].value

	return strconv.Itoa(result), nil
}

type operand struct {
	value   int
	hasSign bool
	posEnd  int
}

func getOperand(runes []rune, start int) (operand, error) {
	var result operand
	var negative bool

	pos := getNextNonSpace(runes, start)
	if pos < 0 {
		return operand{}, noMoreOperands
	}

	if r := runes[pos]; r == '+' {
		result.hasSign = true
		pos += 1
	} else if r == '-' {
		result.hasSign = true
		negative = true
		pos += 1
	}

	result.posEnd = findOperandEnd(runes, pos)

	opStr := string(runes[pos:result.posEnd])
	opStr = strings.TrimSpace(opStr)

	value, err := strconv.Atoi(opStr)
	if err != nil {
		return operand{}, fmt.Errorf("bad operand %q: %w", opStr, err)
	}
	if negative {
		value = -value
	}
	result.value = value
	return result, nil
}

// getNextNonSpace returns the position of the next non-whitespace rune or -1 if the end of slice is reached
func getNextNonSpace(runes []rune, start int) int {
	for i := start; i < len(runes); i++ {
		if !unicode.IsSpace(runes[i]) {
			return i
		}
	}
	return -1
}

// findOperandEnd returns the position of the first symbol of the next operand or len(runes)
func findOperandEnd(runes []rune, start int) int {
	for i := start; i < len(runes); i++ {
		switch runes[i] {
		case '+', '-':
			return i
		}
	}
	return len(runes)
}
