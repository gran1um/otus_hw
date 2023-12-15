package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var result strings.Builder
	var prevRune rune
	escapeMode := false

	for _, r := range s {
		if escapeMode {
			result.WriteRune(r)
			prevRune = r
			escapeMode = false
			continue
		}

		if r == '\\' {
			escapeMode = true
			continue
		}

		if unicode.IsDigit(r) {
			if prevRune == 0 {
				return "", ErrInvalidString
			}
			if err := processDigit(r, prevRune, &result); err != nil {
				return "", err
			}
			continue
		}

		result.WriteRune(r)
		prevRune = r
	}

	if escapeMode {
		return "", ErrInvalidString
	}

	return result.String(), nil
}

func processDigit(digitRune rune, prevRune rune, result *strings.Builder) error {
	count, err := strconv.Atoi(string(digitRune))
	if err != nil {
		return err
	}

	if count == 0 && result.Len() > 0 {
		str := result.String()
		result.Reset()
		result.WriteString(str[:len(str)-1])
	} else if count > 1 {
		result.WriteString(strings.Repeat(string(prevRune), count-1))
	}

	return nil
}
