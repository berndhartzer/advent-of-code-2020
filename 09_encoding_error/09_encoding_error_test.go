package encoding_error

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestEncodingErrorPartOne(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		example := []int{
			35,
			20,
			15,
			25,
			47,
			40,
			62,
			55,
			65,
			95,
			102,
			117,
			150,
			182,
			127,
			219,
			299,
			277,
			309,
			576,
		}
		preamble := 5

		expected := 127
		got, i := EncodingErrorPartOne(example, preamble)
		if got != expected {
			t.Fail()
		}
		if i != 14 {
			t.Fail()
		}
	})

	t.Run("actual input", func(t *testing.T) {
		input := readFileAsIntSlice(t, "./09_input.txt")
		preamble := 25

		output, _ := EncodingErrorPartOne(input, preamble)
		t.Log(fmt.Sprintf("%d\n", output))
	})
}

func TestEncodingErrorPartTwo(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		example := []int{
			35,
			20,
			15,
			25,
			47,
			40,
			62,
			55,
			65,
			95,
			102,
			117,
			150,
			182,
			127,
			219,
			299,
			277,
			309,
			576,
		}
		preamble := 5

		expected := 62
		got := EncodingErrorPartTwo(example, preamble)
		if got != expected {
			t.Fail()
		}
	})

	t.Run("actual input", func(t *testing.T) {
		input := readFileAsIntSlice(t, "./09_input.txt")
		preamble := 25

		output := EncodingErrorPartTwo(input, preamble)
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
