package crab_combat

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestCrabCombatPartOne(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		example := []string{
			"Player 1:",
			"9",
			"2",
			"6",
			"3",
			"1",
			"",
			"Player 2:",
			"5",
			"8",
			"4",
			"7",
			"10",
		}

		expected := 306
		got := CrabCombatPartOne(example)
		if got != expected {
			t.Fail()
		}
	})

	t.Run("actual input", func(t *testing.T) {
		input := readFileAsStringSlice(t, "./input.txt")

		output := CrabCombatPartOne(input)
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
