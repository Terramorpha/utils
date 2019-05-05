package utils

func ParseLine(s string, separator rune) []string {
	var (
		outputSlice = make([]string, 0)
		currentWord []rune

		doubleQuoted bool
		quoted       bool
		escaped      bool
	)
	currentWord = make([]rune, 0, 32)
	for _, char := range s { //ranging through each char
		//fmt.Printf("char: %c\n", char)
		//fmt.Println(string(currentWord))
		if escaped { // if char before was \ :
			currentWord = append(currentWord, char)
			escaped = false
			continue
		}

		if char == '\\' {
			escaped = true
			continue
		}

		if quoted {
			if char == '\'' {
				quoted = false
				continue
			}
			currentWord = append(currentWord, char)
			continue
		}

		if doubleQuoted {
			if char == '"' {
				doubleQuoted = false
				continue
			}
			currentWord = append(currentWord, char)
			continue
		}

		if char == '\'' {
			quoted = true
			continue
		}

		if char == '"' {
			doubleQuoted = true
			continue
		}
		if char == separator {
			outputSlice = append(outputSlice, string(currentWord))
			currentWord = make([]rune, 0, 32)
			//fmt.Println(outputSlice)
			continue
		}
		currentWord = append(currentWord, char)

	}
	if len(currentWord) > 0 {
		outputSlice = append(outputSlice, string(currentWord))
	}

	return outputSlice
}

func removeEmpty(in []string) (out []string) {
	for i := range in {
		if in[i] != "" {
			out = append(out, in[i])
		}
	}
	return
}
