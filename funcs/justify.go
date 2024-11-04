// package funcs

// import (
// 	"fmt"
// 	"os"
// 	"strings"

// 	"golang.org/x/term"
// )

// func PrintJustify(sentences []string, banner []string, align string, w int) {
// 	asciiLines := make([]string, 9)

// 	width, _, err := term.GetSize(int(os.Stdout.Fd()))
// 	if err != nil {
// 		fmt.Println("Error getting terminal size:", err)
// 		return
// 	}

// 	fmt.Println("Max terminal width:", width)

// 	//to fit with different print sizes
// 	for h := 1; h < 9; h++ {
// 		for _, asciiLine := range asciiLines {
// 			padding := ""
// 			switch align {
// 			case "justify":

// 			case "center":
// 				padding = strings.Repeat(" ", (width-len(asciiLine))/2)
// 			case "right":
// 				padding = strings.Repeat(" ", (width - len(asciiLine)))
// 			case "left":
// 				padding = ""
// 			default:
// 				padding = ""
// 			}
// 			fmt.Println(padding + asciiLine)
// 		}
// 		// fmt.Println(writer)
// 	}
// }

package funcs

import (
	"fmt"
	"strings"
)

func PrintJustify(sentences []string, banner []string, align string, width int) {
	// asciiLines := make([]string, 9) // To hold lines of ASCII art for each character height

	for _, sentence := range sentences {
		if sentence == "" {
			fmt.Println() // Blank line for empty sentences
			continue
		}

		// Convert each character in the sentence into ASCII art lines
		for h := 1; h < 9; h++ { // ASCII art character height is 8
			lineBuilder := []string{} // To build the justified line
			lineWidth := 0            // Keep track of current line width

			for _, char := range sentence {
				if char < 32 || char > 126 { // Skip non-printable characters
					continue
				}

				// Each character in banner has 9 lines, calculate the offset
				lineIndex := (int(char) - 32) * 9 // Find the starting line index for the character in banner
				line := banner[lineIndex+h]       // Get the corresponding ASCII art line for character height

				lineBuilder = append(lineBuilder, line) // Append character line
				lineWidth += len(line)                  // Increment total width
			}

			// Now handle alignment based on `align` type
			switch align {
			case "justify":
				// Calculate spaces to distribute for justification
				if lineWidth < width && len(lineBuilder) > 0 {
					spaceNeeded := width - lineWidth                    // Total spaces required to fill line
					spacesBetweenWords := len(lineBuilder) - 1          // Number of spaces between words
					extraSpaces := spaceNeeded / spacesBetweenWords     // Base spaces to add between words
					remainderSpaces := spaceNeeded % spacesBetweenWords // Extra spaces to distribute

					// Rebuild line with evenly distributed spaces
					justifiedLine := ""
					for i, segment := range lineBuilder {
						justifiedLine += segment
						if i < spacesBetweenWords {
							justifiedLine += strings.Repeat(" ", extraSpaces) // Add base spaces
							if i < remainderSpaces {                          // Add remainder spaces
								justifiedLine += " "
							}
						}
					}
					fmt.Println(justifiedLine) // Print the justified line
				} else {
					// If line width >= terminal width or only one word, print left-aligned
					fmt.Println(strings.Join(lineBuilder, ""))
				}
			case "center":
				padding := strings.Repeat(" ", (width-lineWidth)/2)
				fmt.Println(padding + strings.Join(lineBuilder, ""))
			case "right":
				padding := strings.Repeat(" ", width-lineWidth)
				fmt.Println(padding + strings.Join(lineBuilder, ""))
			case "left":
				fmt.Println(strings.Join(lineBuilder, " "))
			default:
				fmt.Println(strings.Join(lineBuilder, " "))
			}
		}
	}
}
