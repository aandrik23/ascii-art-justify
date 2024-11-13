package main

import (
	"fmt"
	"justify/funcs"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func getTerminalWidth() (int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("could not get terminal width")
	}

	// Output format: "<rows> <columns>"
	parts := strings.Fields(string(out))
	if len(parts) < 2 {
		return 0, fmt.Errorf("unexpected output from stty size")
	}

	width, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("failed to parse terminal width")
	}
	return width, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("ERROR! Usage: go run . [OPTION] [STRING] [BANNER] \nEX: go run . --align=<TYPE> something standard ")
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
			fmt.Printf("ERROR! Usage: go run . [OPTION] [STRING] [BANNER] \nEX: go run . --align=%v something standard \n", os.Args[2])
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

	if len(sepText) == 0 {
		fmt.Println("The text is nil")
		return
	}

	file, err := os.ReadFile("banners/" + styleBanner + ".txt")
	if err != nil {
		fmt.Println("banner does not exist.")
		return
	}

	width, err := getTerminalWidth()
	if err != nil {
		fmt.Println("Error getting terminal width:", err)
		return
	}

	str := string(file)
	str = strings.Replace(str, "\r\n", "\n", -1)
	lines := strings.Split(str, "\n")
	if !hasAlign {
		funcs.PrintAsciiArt(sepText, lines)
	} else if align == "justify" {
		funcs.PrintJustify(sepText, lines, align, width)
	} else {
		funcs.PrintAlign(sepText, lines, align, width)
	}
}
