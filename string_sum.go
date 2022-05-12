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
	fmt.Printf("input: %q\n", input)
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
	value    int
	hasSign  bool
	posStart int
	posEnd   int
}

func getOperand(runes []rune, start int) (operand, error) {
	var result operand
	var negative bool

	for i := start; i < len(runes); i++ {
		r := runes[i]
		if r == '+' {
			if result.hasSign {
				return operand{}, fmt.Errorf("multiple operators")
			}
			result.hasSign = true
			continue
		}
		if r == '-' {
			if result.hasSign {
				return operand{}, fmt.Errorf("multiple operators")
			}
			result.hasSign = true
			negative = true
			continue
		}
		if unicode.IsDigit(r) {
			result.posStart = i
			valRunes := getValueRunes(runes, i)
			value, err := strconv.Atoi(string(valRunes))
			if err != nil {
				return operand{}, fmt.Errorf("bad operand value: %w", err)
			}
			result.value = value
			if negative {
				result.value = -result.value
			}
			result.posEnd = result.posStart + len(valRunes)
			return result, nil
		}
		if unicode.IsSpace(r) {
			continue
		}
		return operand{}, fmt.Errorf("bad symbol %c", r)
	}

	return operand{}, noMoreOperands
}

func getValueRunes(runes []rune, start int) []rune {
	i := start + 1
	for ; i < len(runes); i++ {
		r := runes[i]
		if !unicode.IsDigit(r) {
			break
		}
	}
	return runes[start:i]
}
