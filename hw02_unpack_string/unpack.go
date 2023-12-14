package hw02unpackstring

import (
	"errors"
	"strconv"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if len(s) > 0 && unicode.IsDigit(rune(s[0])) {
		return "", ErrInvalidString
	}

	var result string
	var prevRune rune
	escapeMode := false

	for i, r := range s {
		if !unicode.IsDigit(r) && i+1 < len(s) && unicode.IsDigit(rune(s[i+1])) && rune(s[i+1]) == '0' {
			runes := []rune(result)
			result = string(runes[:len(runes)-1])
			i += 2
		}

		if escapeMode {
			if r != '\\' && !unicode.IsDigit(r) {
				return "", errors.New("invalid escape sequence")
			}
			if i+1 < len(s) && unicode.IsDigit(rune(s[i+1])) {
				prevRune = r
			}
			result += string(r)
			escapeMode = false
			continue
		}

		if unicode.IsDigit(r) && i+1 < len(s) && unicode.IsDigit(rune(s[i+1])) {
			return "", ErrInvalidString
		}

		if r == '\\' {
			escapeMode = true
			continue
		}

		if unicode.IsDigit(r) {
			if prevRune == 0 {
				return "", errors.New("string starts with a digit")
			}
			count, _ := strconv.Atoi(string(r))
			for i := 0; i < count-1; i++ {
				result += string(prevRune)
			}
		} else {
			result += string(r)
		}
		prevRune = r
	}

	return result, nil
}
