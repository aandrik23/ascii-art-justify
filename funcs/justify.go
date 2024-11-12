// package funcs

// import (
// 	"fmt"
// 	"strings"
// )

// func PrintJustify(sentences []string, banner []string, align string, width int) {
// 	asciiSpace := []string{
// 		"     ", // Adjust width here if necessary
// 		"     ",
// 		"     ",
// 		"     ",
// 		"     ",
// 		"     ",
// 		"     ",
// 		"     ",
// 	}
// 	for i := 0; i < 8; i++ {
// 		asciiSpace[i] = "  "
// 	}
// 	for _, sentence := range sentences {
// 		if sentence == "" {
// 			fmt.Println() // Blank line for empty sentences
// 			continue
// 		}

// 		// Split sentence into words to handle justification between words
// 		words := strings.Fields(sentence) // Split on whitespace

// 		// Collect ASCII lines for each word separately
// 		asciiLines := make([]string, 8) // 8 lines per ASCII character height

// 		for h := 1; h <= 8; h++ { // ASCII art character height is 8
// 			wordLines := []string{}
// 			lineWidth := 0

// 			// Build the ASCII art for each word in the sentence
// 			for _, word := range words {
// 				wordAscii := ""
// 				for _, char := range word {
// 					if char < 32 || char > 126 {
// 						continue
// 					}

// 					// Calculate the line index in the banner for each character
// 					lineIndex := (int(char) - 32) * 9
// 					wordAscii += banner[lineIndex+h] // Add the line for this ASCII character
// 				}
// 				wordLines = append(wordLines, wordAscii)
// 				lineWidth += len(wordAscii)
// 			}

// 			// Handle alignment based on the `align` type
// 			switch align {
// 			case "justify":
// 				if lineWidth < width && len(wordLines) > 1 { // Justify if line width is less than target and multiple words
// 					spaceNeeded := width - lineWidth
// 					spacesBetweenWords := len(wordLines) - 1
// 					extraSpaces := spaceNeeded / spacesBetweenWords
// 					remainderSpaces := spaceNeeded % spacesBetweenWords

// 					// Build the justified line with spaces only between words
// 					justifiedLine := ""
// 					for i, wordLine := range wordLines {
// 						justifiedLine += wordLine
// 						if i < spacesBetweenWords { // Add spaces between words only
// 							justifiedLine += strings.Repeat(" ", extraSpaces)
// 							if i < remainderSpaces { // Distribute remaining spaces
// 								justifiedLine += " "
// 							}
// 						}
// 					}
// 					asciiLines[h-1] = justifiedLine
// 				} else {
// 					asciiLines[h-1] = strings.Join(wordLines, asciiSpace[h-1]) // Left-align if width is met or single word
// 				}
// 			case "center":
// 				padding := strings.Repeat(" ", (width-lineWidth)/2)
// 				asciiLines[h-1] = padding + strings.Join(wordLines, asciiSpace[h-1])
// 			case "right":
// 				padding := strings.Repeat(" ", width-(lineWidth+len(words)-1))

// 				asciiLines[h-1] = padding + strings.Join(wordLines, " ")
// 			case "left":
// 				asciiLines[h-1] = strings.Join(wordLines, asciiSpace[h-1])
// 			default:
// 				asciiLines[h-1] = strings.Join(wordLines, asciiSpace[h-1])

// 			}
// 		}

// 		// Print each line of the final justified ASCII art for the sentence
// 		for _, line := range asciiLines {
// 			fmt.Println(line)
// 		}
// 	}
// }

package funcs

import (
	"fmt"
	"strings"
)

func PrintJustify(sentences []string, banner []string, align string, width int) {
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

		for h := 1; h <= 8; h++ { // ASCII art character height is 8
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
			case "justify":
				if lineWidth < width && len(wordLines) > 1 {
					// Justify if line width is less than target and multiple words
					spaceNeeded := width - lineWidth
					spacesBetweenWords := len(wordLines) - 1
					extraSpaces := spaceNeeded / spacesBetweenWords
					remainderSpaces := spaceNeeded % spacesBetweenWords

					// Build the justified line by adding extra spaces between words
					justifiedLine := ""
					for i, wordAscii := range wordLines {
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
					asciiLines[h-1] = strings.Join(wordLines, "")
				}
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
