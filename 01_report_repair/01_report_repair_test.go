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
		file, _ := os.Open("./01_input.txt")
		defer file.Close()

		var input []int
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			num, _ := strconv.Atoi(line)
			input = append(input, num)
		}

		output := ReportRepairPartOne(input)
		t.Log(fmt.Sprintf("%d\n", output))
	})
}
