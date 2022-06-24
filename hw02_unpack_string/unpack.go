package hw02unpackstring

import (
	"errors"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var result string
	strLength := utf8.RuneCountInString(str)
	for i := 0; i < strLength; i++ {
		curr := rune(str[i])
		if i == 0 && unicode.IsNumber(curr) {
			return "", ErrInvalidString
		}

		if !unicode.IsNumber(curr) {
			if i != strLength-1 && unicode.IsNumber(rune(str[i+1])) {
				for j := 0; j < int(rune((str[i+1])-'0')); j++ {
					result += string(curr)
				}
			} else {
				result += string(curr)
			}
		} else if i != strLength-1 && unicode.IsNumber(rune(str[i+1])) {
			return "", ErrInvalidString
		}
	}
	return result, nil
}
