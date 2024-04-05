package cmd

import "strings"

var romanNumerals = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func ConvertArabicToRoman(number int) string {
	if number == 0 {
		return
	}
	var buffer strings.Builder
	for value, symbol := range romanNumerals {
		for number >= value {
			buffer.WriteString(symbol)
			number -= value
		}
	}
	return buffer.String()
}
