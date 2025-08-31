package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func slashAdd(slashStr string) (string, int, error) {
	if len(slashStr) < 2 {
		return "", 0, ErrInvalidString
	}

	ch := slashStr[1]
	if len(slashStr) != 2 && unicode.IsDigit(rune(slashStr[2])) {
		return strings.Repeat(string(ch), int(slashStr[2]-'0')), 2, nil
	}

	return string(ch), 1, nil
}

func Unpack(str string) (string, error) {
	var finStr strings.Builder
	var ch, s rune
	var valid bool
	for i := 0; i < len(str); i++ {
		s = rune(str[i])
		if !valid && unicode.IsDigit(s) {
			return "", ErrInvalidString
		}

		if !unicode.IsDigit(s) {
			if valid {
				finStr.WriteRune(ch)
			}
			ch = s
			valid = true
		}

		if s == '\\' {
			slashStr, skip, err := slashAdd(str[i:])
			if err != nil {
				return "", err
			}

			finStr.WriteString(slashStr)
			valid = false
			i += skip
			continue
		}

		if valid && unicode.IsDigit(s) {
			finStr.WriteString(strings.Repeat(string(ch), int(s-'0')))
			valid = false
		}
	}

	if valid {
		finStr.WriteRune(ch)
	}

	return finStr.String(), nil
}
