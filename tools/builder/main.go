package main

import (
	"fmt"
	"strings"
)

const (
	h1Prefix = "# "
	h2Prefix = "## "
	h3Prefix = "### "

	boldWrap = "**"
)

func ParseHeading(input string) (string, error) {
	if strings.HasPrefix(input, h1Prefix) {
		return parseH1(input), nil
	}

	if strings.HasPrefix(input, h2Prefix) {
		return parseH2(input), nil
	}

	if strings.HasPrefix(input, h3Prefix) {
		return parseH3(input), nil
	}

	return "", fmt.Errorf(InvalidSyntax)
}

func ParseBold(input string) (string, error) {
	if !strings.HasPrefix(input, boldWrap) || !strings.HasSuffix(input, boldWrap) {
		return "", fmt.Errorf(InvalidSyntax)
	}

	text, _ := strings.CutPrefix(input, boldWrap)
	text, _ = strings.CutSuffix(text, boldWrap)
	return fmt.Sprintf("<b>%s</b>", text), nil
}

func parseH1(input string) string {
	text, _ := strings.CutPrefix(input, h1Prefix)
	return fmt.Sprintf("<h1>%s</h1>", text)
}

func parseH2(input string) string {
	text, _ := strings.CutPrefix(input, h2Prefix)
	return fmt.Sprintf("<h2>%s</h2>", text)
}

func parseH3(input string) string {
	text, _ := strings.CutPrefix(input, h3Prefix)
	return fmt.Sprintf("<h3>%s</h3>", text)
}

func Parse(input string) (string, error) {
	text, err := ParseHeading(input)
	if err != nil {
		return "", err
	}
	text, err = ParseBold(text)
	if err != nil {
		return "", err
	}
	return text, nil
}
