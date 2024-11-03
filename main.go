package main

import (
	"fmt"
	"justify/funcs"
	"os"
	"strings"

	"golang.org/x/term"
)

func main() {
	// Check if the number of command-line arguments is correct
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("ERROR! Usage: go run . [OPTION] [STRING] [BANNER] \nEX: go run . --align=<type> something standard")
		return
	}
	var text string
	var align string
	styleBanner := "standard"
	hasAlign := false

	if len(os.Args) == 4 {
		align = os.Args[1]
		text = os.Args[2]
		styleBanner = strings.ToLower(os.Args[3])
		hasAlign = true

		if !strings.HasPrefix(align, "--align=") {
			fmt.Println("ERROR! Usage: go run . [OPTION] [STRING] [BANNER] \nEX: go run . --align=<type> something standard")
			return
		}

		align = align[8:] // Skips the "--align=" part
	}

	if len(os.Args) == 3 {
		align = os.Args[1]
		if strings.HasPrefix(align, "--align=") {

			text = os.Args[2]
			align = align[8:] // Skips the "--align=" part
			hasAlign = true
		} else {

			text = os.Args[1]
			styleBanner = strings.ToLower(os.Args[2])
		}
	}

	if len(os.Args) == 2 {
		text = os.Args[1]
	}

	sepText := strings.Split(text, "\\n")

	file, err := os.ReadFile("banners/" + styleBanner + ".txt")
	if err != nil {
		fmt.Println(" banner does not exist.")
		return
	}

	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

	fmt.Println("Max terminal width:", width)

	str := string(file)
	str = strings.Replace(str, "\r\n", "\n", -1)
	lines := strings.Split(str, "\n")
	if !hasAlign {

		funcs.PrintAsciiArt(sepText, lines)
	} else {

		funcs.PrintJustify(sepText, lines, align, width)
	}
}
