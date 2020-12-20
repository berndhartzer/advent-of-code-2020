package operation_order

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestOperationOrderPartOne(t *testing.T) {
	tests := map[string]struct {
		input    []string
		expected int
	}{
		"example 1": {input: []string{"1 + 2 * 3 + 4 * 5 + 6"}, expected: 71},
		"example 2": {input: []string{"1 + (2 * 3) + (4 * (5 + 6))"}, expected: 51},
		"example 3": {input: []string{"2 * 3 + (4 * 5)"}, expected: 26},
		"example 4": {input: []string{"5 + (8 * 3 + 9 + 3 * 4 * 3)"}, expected: 437},
		"example 5": {input: []string{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"}, expected: 12240},
		"example 6": {input: []string{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"}, expected: 13632},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if OperationOrderPartOne(tc.input) != tc.expected {
				t.Fail()
			}
		})
	}

	t.Run("actual input", func(t *testing.T) {
		input := readFileAsStringSlice(t, "./input.txt")

		output := OperationOrderPartOne(input)
		t.Log(fmt.Sprintf("%d\n", output))
	})
}

func readFileAsStringSlice(t *testing.T, name string) []string {
	t.Helper()

	file, _ := os.Open(name)
	defer file.Close()

	var contents []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	return contents
}
