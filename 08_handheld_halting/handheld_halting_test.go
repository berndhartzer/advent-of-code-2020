package handheld_halting

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestHandheldHaltingPartOne(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		example := []string{
			"nop +0",
			"acc +1",
			"jmp +4",
			"acc +3",
			"jmp -3",
			"acc -99",
			"acc +1",
			"jmp -4",
			"acc +6",
		}

		expected := 5
		got := HandheldHaltingPartOne(example)
		if got != expected {
			t.Fail()
		}
	})

	t.Run("actual input", func(t *testing.T) {
		input := readFileAsStringSlice(t, "./08_input.txt")

		output := HandheldHaltingPartOne(input)
		t.Log(fmt.Sprintf("%d\n", output))
	})
}

func TestHandheldHaltingPartTwo(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		example := []string{
			"nop +0",
			"acc +1",
			"jmp +4",
			"acc +3",
			"jmp -3",
			"acc -99",
			"acc +1",
			"jmp -4",
			"acc +6",
		}

		expected := 8
		got := HandheldHaltingPartTwo(example)
		if got != expected {
			t.Fail()
		}
	})

	t.Run("actual input", func(t *testing.T) {
		input := readFileAsStringSlice(t, "./08_input.txt")

		output := HandheldHaltingPartTwo(input)
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
