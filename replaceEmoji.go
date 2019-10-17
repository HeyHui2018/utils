package utils

import (
	"unicode/utf8"
)

func ReplaceEmoji(originalStr, replaceStr string) string {
	var newRune []rune
	for _, value := range originalStr {
		_, size := utf8.DecodeRuneInString(string(value))
		if size > 3 {
			newRune = append(newRune, []rune(replaceStr)...)
		} else {
			newRune = append(newRune, value)
		}
	}
	return string(newRune)
}

/*
	if you want to remove emoji,just set replaceStr = "".
	for example:
	new := FilterEmoji("😭123😭asd😄Winnie", "")
*/
