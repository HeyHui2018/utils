package utils

import (
	"regexp"
)

func HaveLetter(s string) (bool, error) {
	ok, err := regexp.MatchString(`.*[A-Za-z]+.*`, s)
	return ok, err
}
