package utils

import (
	"unicode/utf8"
)

//😭123😭asd😄Win惠nie

func FilterEmoji(originalStr, replaceStr string) string {
	var newRune []rune
	for _, value := range originalStr {
		r, size := utf8.DecodeRuneInString(string(value))
		if size > 3 {
			newRune = append(newRune, []rune(replaceStr)...)
		} else {
			newRune = append(newRune, value)
		}
	}
	return string(newRune)
}

/*
	new := FilterEmoji("😭123😭asd😄Win惠nie", "你好")
 */
