package password_philosophy

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestPasswordPhilosophyPartOne(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		example := []string{
			"1-3 a: abcde",
			"1-3 b: cdefg",
			"2-9 c: ccccccccc",
		}
		expected := 2
		got := PasswordPhilosophyPartOne(example)
		if got != expected {
			t.Fail()
		}
	})

	t.Run("actual input", func(t *testing.T) {
		input := readFileAsStringSlice(t, "./02_input.txt")

		output := PasswordPhilosophyPartOne(input)
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
