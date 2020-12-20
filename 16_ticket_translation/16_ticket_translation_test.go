package ticket_translation

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestTicketTranslationPartOne(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		example := []string{
			"class: 1-3 or 5-7",
			"row: 6-11 or 33-44",
			"seat: 13-40 or 45-50",
			"",
			"your ticket:",
			"7,1,14",
			"",
			"nearby tickets:",
			"7,3,47",
			"40,4,50",
			"55,2,20",
			"38,6,12",
		}

		expected := 71
		got := TicketTranslationPartOne(example)
		if got != expected {
			t.Fail()
		}
	})

	t.Run("actual input", func(t *testing.T) {
		input := readFileAsStringSlice(t, "./input.txt")

		output := TicketTranslationPartOne(input)
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
