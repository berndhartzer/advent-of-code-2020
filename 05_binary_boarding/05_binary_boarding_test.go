package binary_boarding

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestBinaryBoardingPartOne(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		example := []string{
			"FBFBBFFRLR", // 357
			"BFFFBBFRRR", // 567
			"FFFBBBFRRR", // 119
			"BBFFBBFRLL", // 820
		}

		var expected int64 = 820
		got := BinaryBoardingPartOne(example)
		if got != expected {
			t.Fail()
		}
	})

	t.Run("actual input", func(t *testing.T) {
		input := readFileAsStringSlice(t, "./05_input.txt")

		output := BinaryBoardingPartOne(input)
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
