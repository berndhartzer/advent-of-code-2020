package rambunctious_recitation

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestRambunctiousRecitationPartOne(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		example := []int{
			0,
			3,
			6,
		}
		limit := 2020

		expected := 436
		got := RambunctiousRecitation(example, limit)
		if got != expected {
			t.Fail()
		}
	})

	t.Run("actual input", func(t *testing.T) {
		input := readFileAsIntSlice(t, "./input.txt")
		limit := 2020

		output := RambunctiousRecitation(input, limit)
		t.Log(fmt.Sprintf("%d\n", output))
	})
}

func TestRambunctiousRecitationPartTwo(t *testing.T) {
	t.Run("actual input", func(t *testing.T) {
		input := readFileAsIntSlice(t, "./input.txt")
		limit := 30000000

		output := RambunctiousRecitation(input, limit)
		t.Log(fmt.Sprintf("%d\n", output))
	})
}

func readFileAsIntSlice(t *testing.T, name string) []int {
	t.Helper()

	file, _ := os.Open(name)
	defer file.Close()

	var contents []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ",")
		for _, n := range split {
			num, _ := strconv.Atoi(n)
			contents = append(contents, num)
		}
	}

	return contents
}
