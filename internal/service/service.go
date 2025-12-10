package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func DetectAndConvert(input string) (string, error) {
	if input == "" {
		return "", nil
	}

	hasMorseChars := strings.ContainsAny(input, ".- ")

	if hasMorseChars {
		result := morse.ToText(input)
		if result == "" {
			return "", errors.New("invalid morse code")
		}
		return result, nil
	} else {
		result := morse.ToMorse(input)
		if result == "" {
			return "", errors.New("conversion to morse failed")
		}
		return result, nil
	}
}
