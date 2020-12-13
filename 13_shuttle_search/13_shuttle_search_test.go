package shuttle_search

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestShuttleSearchPartOne(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		earliest := 939
		example := []string{
			"7",
			"13",
			"x",
			"x",
			"59",
			"x",
			"31",
			"19",
		}

		expected := 295
		got := ShuttleSearchPartOne(earliest, example)
		if got != expected {
			t.Fail()
		}
	})

	t.Run("actual input", func(t *testing.T) {
		input := readFileAsStringSlice(t, "./13_input.txt")

		earliest, _ := strconv.Atoi(input[0])
		buses := strings.Split(input[1], ",")

		output := ShuttleSearchPartOne(earliest, buses)
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
