// package funcs

// import (
// 	"fmt"
// 	"strings"
// )

// func PrintJustify(sentences []string, banner []string, align string, width int) {
// 	// asciiLines := make([]string, 9) // To hold lines of ASCII art for each character height

// 	for _, sentence := range sentences {
// 		if sentence == "" {
// 			fmt.Println() // Blank line for empty sentences
// 			continue
// 		}

// 		// Convert each character in the sentence into ASCII art lines
// 		for h := 1; h < 9; h++ { // ASCII art character height is 8
// 			lineBuilder := []string{} // To build the justified line
// 			lineWidth := 0            // Keep track of current line width

// 			for _, char := range sentence {
// 				if char < 32 || char > 126 { // Skip non-printable characters
// 					continue
// 				}

// 				// Each character in banner has 9 lines, calculate the offset
// 				lineIndex := (int(char) - 32) * 9 // Find the starting line index for the character in banner
// 				line := banner[lineIndex+h]       // Get the corresponding ASCII art line for character height

// 				lineBuilder = append(lineBuilder, line) // Append character line
// 				lineWidth += len(line)                  // Increment total width
// 			}

// 			// Now handle alignment based on `align` type
// 			switch align {
// 			case "justify":
// 				// Calculate spaces to distribute for justification
// 				if width > lineWidth && len(lineBuilder) > 1 {
// 					spaceNeeded := width - lineWidth                    // Total spaces required to fill line
// 					spacesBetweenWords := len(lineBuilder) - 1          // Number of spaces between words
// 					extraSpaces := spaceNeeded / spacesBetweenWords     // Base spaces to add between words
// 					remainderSpaces := spaceNeeded % spacesBetweenWords // Extra spaces to distribute

// 					// Rebuild line with evenly distributed spaces
// 					justifiedLine := ""
// 					for i, segment := range lineBuilder {
// 						justifiedLine += segment
// 						if i < spacesBetweenWords {
// 							justifiedLine += strings.Repeat(" ", extraSpaces) // Add base spaces
// 							if i < remainderSpaces {                          // Add remainder spaces
// 								justifiedLine += " "
// 							}
// 						}
// 					}

// 					fmt.Println(justifiedLine) // Print the justified line
// 				} else {
// 					// If line width >= terminal width or only one word, print left-aligned
// 					fmt.Println(strings.Join(lineBuilder, ""))
// 				}
// 			case "center":
// 				padding := strings.Repeat(" ", (width-lineWidth)/2)
// 				fmt.Println(padding + strings.Join(lineBuilder, ""))
// 			case "right":
// 				padding := strings.Repeat(" ", width-lineWidth)
// 				fmt.Println(padding + strings.Join(lineBuilder, ""))
// 			case "left":
// 				fmt.Println(strings.Join(lineBuilder, " "))
// 			default:
// 				fmt.Println(strings.Join(lineBuilder, " "))
// 			}
// 		}
// 	}
// }

package funcs

import (
	"fmt"
	"strings"
)

func PrintJustify(sentences []string, banner []string, align string, width int) {
	for _, sentence := range sentences {
		if sentence == "" {
			fmt.Println() // Blank line for empty sentences
			continue
		}

		// Split sentence into words to handle justification between words
		words := strings.Fields(sentence) // Split on whitespace

		// Collect ASCII lines for each word separately
		asciiLines := make([]string, 8) // 8 lines per ASCII character height

		for h := 1; h <= 8; h++ { // ASCII art character height is 8
			wordLines := []string{}
			lineWidth := 0

			// Build the ASCII art for each word in the sentence
			for _, word := range words {
				wordAscii := ""
				for _, char := range word {
					if char < 32 || char > 126 {
						continue
					}

					// Calculate the line index in the banner for each character
					lineIndex := (int(char) - 32) * 9
					wordAscii += banner[lineIndex+h] // Add the line for this ASCII character
				}
				wordLines = append(wordLines, wordAscii)
				lineWidth += len(wordAscii)
			}

			// Handle alignment based on the `align` type
			switch align {
			case "justify":
				if lineWidth < width && len(wordLines) > 1 { // Justify if line width is less than target and multiple words
					spaceNeeded := width - lineWidth
					spacesBetweenWords := len(wordLines) - 1
					extraSpaces := spaceNeeded / spacesBetweenWords
					remainderSpaces := spaceNeeded % spacesBetweenWords

					// Build the justified line with spaces only between words
					justifiedLine := ""
					for i, wordLine := range wordLines {
						justifiedLine += wordLine
						if i < spacesBetweenWords { // Add spaces between words only
							justifiedLine += strings.Repeat(" ", extraSpaces)
							if i < remainderSpaces { // Distribute remaining spaces
								justifiedLine += " "
							}
						}
					}
					asciiLines[h-1] = justifiedLine
				} else {
					asciiLines[h-1] = strings.Join(wordLines, " ") // Left-align if width is met or single word
				}
			case "center":
				padding := strings.Repeat(" ", (width-lineWidth)/2)
				asciiLines[h-1] = padding + strings.Join(wordLines, " ")
			case "right":
				padding := strings.Repeat(" ", width-(lineWidth+len(words)-1))

				asciiLines[h-1] = padding + strings.Join(wordLines, " ")
			case "left":
				asciiLines[h-1] = strings.Join(wordLines, " ")
			default:
				asciiLines[h-1] = strings.Join(wordLines, " ")
			}
		}

		// Print each line of the final justified ASCII art for the sentence
		for _, line := range asciiLines {
			fmt.Println(line)
		}
	}
}
