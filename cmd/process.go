package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func process() (err error) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var output string
		if output, err = processLine(line); err != nil {
			return
		}
		fmt.Println(output)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
	}
	return
}

func processLine(line string) (string, error) {
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

func replaceChars(s string) string {
	s = strings.ReplaceAll(s, "J", "I")
	s = strings.ReplaceAll(s, "U", "V")
	s = strings.ReplaceAll(s, "W", "V")
	return s
}
