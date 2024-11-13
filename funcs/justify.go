package funcs

import (
	"fmt"
	"strings"
)

func PrintAlign(sentences []string, banner []string, align string, width int) {
	// Define ASCII art representation of a space character
	asciiSpace := []string{
		"     ", // Adjust width here if necessary
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
	}
	for _, sentence := range sentences {
		if sentence == "" {
			fmt.Println() // Blank line for empty sentences
			continue
		}
		// Split sentence on individual characters to retain multiple spaces
		chars := strings.Split(sentence, "")
		asciiLines := make([]string, 8) // Collect ASCII lines
		for h := 1; h <= 8; h++ {       // ASCII art character height is 8
			wordLines := []string{} // Collect ASCII words in the line
			lineWidth := 0
			// Build the ASCII art for each character in the sentence
			for _, char := range chars {
				charAscii := ""
				if char == " " {
					charAscii = asciiSpace[h-1]
				} else if char[0] >= 32 && char[0] <= 126 {
					lineIndex := (int(char[0]) - 32) * 9
					charAscii = banner[lineIndex+h]
				}
				wordLines = append(wordLines, charAscii)
				lineWidth += len(charAscii)
			}
			// Handle alignment based on the `align` type
			switch align {
			case "center":
				padding := strings.Repeat(" ", (width-lineWidth)/2)
				asciiLines[h-1] = padding + strings.Join(wordLines, "") + padding
			case "right":
				padding := strings.Repeat(" ", width-lineWidth)
				asciiLines[h-1] = padding + strings.Join(wordLines, "")
			case "left":
				asciiLines[h-1] = strings.Join(wordLines, "")
			default:
				asciiLines[h-1] = strings.Join(wordLines, "")
			}
		}
		// Print each line of the final ASCII art for the sentence
		for _, line := range asciiLines {
			fmt.Println(line)
		}
	}
}

func PrintJustify(sentences []string, banner []string, align string, width int) {
	asciiSpace := []string{
		"     ", // Adjust width here if necessary
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
	}

	for _, sentence := range sentences {
		if sentence == "" {
			fmt.Println() // Blank line for empty sentences
			continue
		}

		// Split sentence into words
		words := strings.Fields(sentence)
		asciiLines := make([]string, 8) // Collect ASCII lines

		for h := 1; h <= 8; h++ { // ASCII art character height is 8
			wordAsciiLines := []string{} // Collect ASCII words in the line
			lineWidth := 0

			// Build the ASCII art for each word in the sentence
			for _, word := range words {
				wordAscii := ""
				for _, char := range word {
					if char >= 32 && char <= 126 {
						lineIndex := (int(char) - 32) * 9
						wordAscii += banner[lineIndex+h]
					}
				}
				wordAsciiLines = append(wordAsciiLines, wordAscii)
				lineWidth += len(wordAscii)
			}

			// Handle alignment based on the `align` type
			switch align {
			case "justify":
				if lineWidth < width && len(wordAsciiLines) > 1 {
					// Justify if line width is less than target and multiple words
					spaceNeeded := width - lineWidth
					spacesBetweenWords := len(wordAsciiLines) - 1
					extraSpaces := spaceNeeded / spacesBetweenWords
					remainderSpaces := spaceNeeded % spacesBetweenWords

					// Build the justified line by adding extra spaces between words
					justifiedLine := ""
					for i, wordAscii := range wordAsciiLines {
						justifiedLine += wordAscii
						if i < spacesBetweenWords { // Add spaces only between words
							justifiedLine += strings.Repeat(" ", extraSpaces)
							if i < remainderSpaces { // Distribute remainder spaces
								justifiedLine += " "
							}
						}
					}
					asciiLines[h-1] = justifiedLine
				} else {
					// Left-align if the width is already met or single word
					asciiLines[h-1] = strings.Join(wordAsciiLines, asciiSpace[h-1])
				}
			}
		}
		for _, line := range asciiLines {
			fmt.Println(line)

		}
	}
}
