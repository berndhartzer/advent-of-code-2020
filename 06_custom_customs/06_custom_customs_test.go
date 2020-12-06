package custom_customs

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestCustomCustomsPartOne(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		example := []string{
			"abc",
			"",
			"a",
			"b",
			"c",
			"",
			"ab",
			"ac",
			"",
			"a",
			"a",
			"a",
			"a",
			"",
			"b",
		}

		expected := 11
		got := CustomCustomsPartOne(example)
		if got != expected {
			t.Fail()
		}
	})

	t.Run("actual input", func(t *testing.T) {
		input := readFileAsStringSlice(t, "./06_input.txt")

		output := CustomCustomsPartOne(input)
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
