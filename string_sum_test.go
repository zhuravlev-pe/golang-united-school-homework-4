package string_sum

import "testing"

func TestStringSum(t *testing.T) {
	data := []struct {
		name     string
		input    string
		result   string
		hasError bool
	}{
		{"simple", "2+3", "5", false},
		{"whitespace", " 2 + 3 ", "5", false},
		{"diff", " 2 - 3 ", "-1", false},
		{"unary minus", " -2 - 3 ", "-5", false},
		{"panic case", "24c+55", "", true},
	}
	for _, datum := range data {
		t.Run(datum.name, func(t *testing.T) {
			result, err := StringSum(datum.input)
			if datum.hasError && err == nil {
				t.Error("error expected")
				return
			}
			if datum.result != result {
				t.Errorf("invalid result. got: %q want: %q", result, datum.result)
			}
		})
	}
}
