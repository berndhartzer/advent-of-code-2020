package rain_risk

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestRainRiskPartOne(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		example := []string{
			"F10",
			"N3",
			"F7",
			"R90",
			"F11",
		}

		expected := 25
		got := RainRiskPartOne(example)
		if got != expected {
			t.Fail()
		}
	})

	t.Run("actual input", func(t *testing.T) {
		input := readFileAsStringSlice(t, "./12_input.txt")

		output := RainRiskPartOne(input)
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
