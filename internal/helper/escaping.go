package helper

import "strings"

func EscapeCharacters(str string) string {
	charsToEscape := []string{
		"\u200B_",
		"\u200B*",
		"\u200B[",
		"\u200B]",
		"\u200B(",
		"\u200B)",
		"\u200B~",
		"\u200B`",
		"\u200B>",
		"\u200B#",
		"\u200B+",
		"\u200B-",
		"\u200B=",
		"\u200B{",
		"\u200B}",
	}
	for _, char := range charsToEscape {
		str = strings.ReplaceAll(str, char, "\\"+char)
	}
	return str
}
