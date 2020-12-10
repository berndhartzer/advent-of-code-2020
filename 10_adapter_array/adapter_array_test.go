package adapter_array

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestAdapterArrayPartOne(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		example := []int{
			16,
			10,
			15,
			5,
			1,
			11,
			7,
			19,
			6,
			12,
			4,
		}

		expected := 35
		got := AdapterArrayPartOne(example)
		if got != expected {
			t.Fail()
		}
	})

	t.Run("example input two", func(t *testing.T) {
		example := []int{
			28,
			33,
			18,
			42,
			31,
			14,
			46,
			20,
			48,
			47,
			24,
			23,
			49,
			45,
			19,
			38,
			39,
			11,
			1,
			32,
			25,
			35,
			8,
			17,
			7,
			9,
			4,
			2,
			34,
			10,
			3,
		}

		expected := 220
		got := AdapterArrayPartOne(example)
		if got != expected {
			t.Fail()
		}
	})

	t.Run("actual input", func(t *testing.T) {
		input := readFileAsIntSlice(t, "./10_input.txt")

		output := AdapterArrayPartOne(input)
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
