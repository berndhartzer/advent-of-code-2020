package report_repair

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestReportRepairPartOne(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		example := []int{1721, 979, 366, 299, 675, 1456}
		expected := 514579
		got := ReportRepairPartOne(example)
		if got != expected {
			t.Fail()
		}
	})

	t.Run("actual input", func(t *testing.T) {
		input := readFileAsIntSlice(t, "./01_input.txt")

		output := ReportRepairPartOne(input)
		t.Log(fmt.Sprintf("%d\n", output))
	})
}

func TestReportRepairPartTwo(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		example := []int{1721, 979, 366, 299, 675, 1456}
		expected := 241861950
		got := ReportRepairPartTwo(example)
		if got != expected {
			t.Fail()
		}
	})

	t.Run("actual input", func(t *testing.T) {
		input := readFileAsIntSlice(t, "./01_input.txt")

		output := ReportRepairPartTwo(input)
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
		num, _ := strconv.Atoi(line)
		contents = append(contents, num)
	}

	return contents
}
