package mbalign

import (
	"math"
)

const (
	LeftAlign uint = 1 << iota
	RightAlign
)

func Align(input string, length int, padchar byte, align uint) string {
	width := 0
	restCount := 0
	var outRunes []rune
	for _, runeValue := range input {
		byteNum := len([]byte(string(runeValue)))

		if byteNum > 1 {
			width += 2
		} else {
			width += 1
		}

		if width <= length {
			outRunes = append(outRunes, runeValue)
		} else {
			break
		}
		restCount = int(math.Abs(float64(length - width)))
	}

	result := string(outRunes)
	padchars := ""
	for i := 0; i < restCount; i++ {
		padchars += string(padchar)
	}

	if align == LeftAlign {
		result = result + padchars
	} else {
		result = padchars + result
	}
	return result
}

func GetWidth(input string) int {
	width := 0
	for _, runeValue := range input {
		byteNum := len([]byte(string(runeValue)))
		if byteNum > 1 {
			width += 2
		} else {
			width += 1
		}
	}

	return width
}
