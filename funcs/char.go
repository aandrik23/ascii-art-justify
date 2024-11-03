package funcs

func getAsciiChar(char rune, banner []string) []string {
	index := (int(char) - 32) * 9
	if index < 1 || index+9 > len(banner) {
		return make([]string, 9)
	}
	return banner[index : index+9]
}
