package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	rom "github.com/brandenc40/romannumeral"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func process() (err error) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		firstpass, err := processLetters(line)
		if err != nil {
			panic(err)
		}
		secondpass, err := processNumerals(firstpass)
		if err != nil {
			panic(err)
		}
		fmt.Println(secondpass)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
	}
	return
}

func processLetters(line string) (string, error) {
	// Strip out accented characters
	transformer := transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool {
		return unicode.Is(unicode.Mn, r)
	}), norm.NFC)
	result, _, _ := transform.String(transformer, line)
	// Strip out punctuation
	var buffer strings.Builder
	for _, char := range result {
		if !unicode.IsPunct(char) {
			buffer.WriteRune(char)
		}
	}
	return replaceChars(strings.ToUpper(buffer.String())), nil
}

func processNumerals(line string) (roman string, err error) {
	re := regexp.MustCompile(`\d+`)
	roman = re.ReplaceAllStringFunc(line,
		func(match string) string {
			num, err := strconv.Atoi(match) // Convert string to integer
			if err != nil {
				fmt.Println(err)
				return ""
			}
			roman, err := rom.IntToString(num)
			if err != nil {
				fmt.Println(err)
			}
			return roman
		})
	return
}

func replaceChars(s string) string {
	s = strings.ReplaceAll(s, "J", "I")
	s = strings.ReplaceAll(s, "U", "V")
	s = strings.ReplaceAll(s, "W", "V")
	return s
}
