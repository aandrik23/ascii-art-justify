package funcs

import (
	"bytes"
	"justify/funcs"
	"os"
	"strings"
	"testing"
)

// Capture output to compare with expected ASCII art
func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	old := os.Stdout
	defer func() {
		os.Stdout = old
	}()
	os.Stdout = w

	f()

	w.Close()

	var buf bytes.Buffer
	buf.ReadFrom(r)

	return buf.String()
}

// Test for character A
func TestPrintJustify_A(t *testing.T) {
	testPrintJustify(t, []string{"A"})
}

// Test for character B
func TestPrintJustify_B(t *testing.T) {
	testPrintJustify(t, []string{"B"})
}

// Test for character C
func TestPrintJustify_C(t *testing.T) {
	testPrintJustify(t, []string{"C"})
}

// Test for multiple characters
func TestPrintJustify_AB(t *testing.T) {
	testPrintJustify(t, []string{"Hello"})
}

// Helper function to test PrintAsciiArt with input and expected output
func testPrintJustify(t *testing.T, input []string) {
	styleBanner := "standard"
	bannerFileName := "../banners/" + styleBanner + ".txt"

	fileContent, err := os.ReadFile(bannerFileName)
	if err != nil {
		t.Fatalf("Failed to read banner file: %v", err)
	}

	fileContentString := strings.ReplaceAll(string(fileContent), "\r\n", "\n")
	bannerLines := strings.Split(fileContentString, "\n")

	expectedOutput := generateExpectedOutput(bannerLines, input)

	align := "left"
	width := 50

	actualOutput := captureOutput(func() {
		funcs.PrintJustify(input, bannerLines, align, width)
	})

	if actualOutput != expectedOutput {
		t.Errorf("For input %v, expected:\n%s\nGot:\n%s", input, expectedOutput, actualOutput)
	}
}

// Helper function to generate expected output based on input
func generateExpectedOutput(bannerLines []string, input []string) string {
	var output strings.Builder

	// Iterate over the height of the ASCII art (8 lines)
	for h := 1; h < 9; h++ {
		for _, word := range input {
			for i := 0; i < len(word); i++ {
				charIndex := int(word[i]) - 32
				lineIndex := charIndex*9 + h
				output.WriteString(bannerLines[lineIndex])
			}
		}
		output.WriteString("\n") // Newline after each row of the characters
	}

	return output.String()
}
