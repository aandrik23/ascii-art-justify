// package funcs

// import (
// 	"fmt"
// 	"os"
// 	"strings"

// 	"golang.org/x/term"
// )

// func PrintJustify(sentences []string, banner []string, align string, w int) {
// 	asciiLines := make([]string, 9)
// 	for _, word := range sentences {

// 	var extraSpaces int
// 		var remainingSpaces int
// 		if numGaps > 0 {
// 			extraSpaces = spaces / numGaps
// 			remainingSpaces = spaces % numGaps
// 		}

// 	width, _, err := term.GetSize(int(os.Stdout.Fd()))
// 	if err != nil {
// 		fmt.Println("Error getting terminal size:", err)
// 		return
// 	}

// 	fmt.Println("Max terminal width:", width)

// 	//to fit with different print sizes
// 	for h := 1; h < 9; h++{
// 	for i, asciiLine := range asciiLines {
// 		padding := ""
// 		switch align {
// 		case "justify":
// 			for j := 0; j < len(asciiLine); j++ {
// 				for lineIndex, line := range banner {
// 					if lineIndex == (int(asciiLine[j])-32)*9+h {
// 						fmt.Print(line)
// 						break
// 					}
// 				}
// 			}

// 			// Print spaces between words if not the last word
// 			if i < numGaps {
// 				for s := 0; s < extraSpaces; s++ {
// 					fmt.Print(" ")
// 				}
// 				if remainingSpaces > 0 {
// 					fmt.Print(" ")
// 					remainingSpaces--
// 				}
// 			}

// 		case "center":
// 			padding = strings.Repeat(" ", (width-len(asciiLine))/2)
// 		case "right":
// 			padding = strings.Repeat(" ", (width - len(asciiLine)))
// 		case "left":
// 			padding = ""
// 		default:
// 			padding = ""
// 		}
// 		fmt.Println(padding + asciiLine)
// 	}
// 	// fmt.Println(writer)
// }

package funcs

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

// PrintJustify prints sentences with ASCII art banner in specified alignment.
func PrintJustify(sentences []string, banner []string, align string, w int) {
	// Get the width of the terminal
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

	fmt.Println("Max terminal width:", width)

	// Prepare ASCII lines for each sentence
	for _, sentence := range sentences {
		asciiLines := make([]string, 9) // Placeholder for 9 lines of ASCII representation

		// For each character in the sentence, accumulate ASCII lines
		for _, char := range sentence {
			if char < 32 || int(char-32)*9+8 >= len(banner) {
				// Skip non-printable ASCII or out-of-range characters
				continue
			}
			startIndex := (int(char) - 32) * 9
			for i := 0; i < 9; i++ {
				asciiLines[i] += banner[startIndex+i]
			}
		}

		// Print each ASCII line with specified alignment
		for _, asciiLine := range asciiLines {
			padding := ""

			// Handle alignment options
			switch align {
			case "justify":
				// Calculate spaces between words
				words := strings.Fields(asciiLine)
				numGaps := len(words) - 1
				spaces := w - len(asciiLine)
				extraSpaces := 0
				remainingSpaces := 0

				if numGaps > 0 {
					extraSpaces = spaces / numGaps
					remainingSpaces = spaces % numGaps
				}

				// Join words with calculated spaces
				justifiedLine := words[0]
				for i := 1; i < len(words); i++ {
					justifiedLine += strings.Repeat(" ", extraSpaces)
					if remainingSpaces > 0 {
						justifiedLine += " "
						remainingSpaces--
					}
					justifiedLine += words[i]
				}
				fmt.Println(justifiedLine)

			case "center":
				padding = strings.Repeat(" ", (width-len(asciiLine))/2)
				fmt.Println(padding + asciiLine + " ")

			case "right":
				padding = strings.Repeat(" ", width-len(asciiLine))
				fmt.Println(padding + asciiLine)

			case "left":
				fmt.Println(asciiLine)

			default:
				fmt.Println(asciiLine) // Default to left alignment
			}
		}
	}
}
