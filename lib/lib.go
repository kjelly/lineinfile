package lib

import (
	"regexp"
	"strings"
)

// Replace replace pattern with text
func Replace(lines []string, re *regexp.Regexp, text string) []string {
	var outputs []string
	outputs = handleLines(lines, re, func(l string, m []string, o []string, re *regexp.Regexp) []string {
		return append(o, strings.Replace(l, m[0], text, -1))

	}, nil)
	return outputs
}

// InsertAfter insert text after pattern.
func InsertAfter(lines []string, re *regexp.Regexp, text string) []string {
	var outputs []string
	outputs = handleLines(lines, re, func(l string, m []string, o []string, re *regexp.Regexp) []string {
		return append(o, strings.Replace(l, m[0], m[0]+text, -1))

	}, nil)
	return outputs
}

// InsertBefore insert text before pattern.
func InsertBefore(lines []string, re *regexp.Regexp, text string) []string {
	var outputs []string
	outputs = handleLines(lines, re, func(l string, m []string, o []string, re *regexp.Regexp) []string {
		return append(o, strings.Replace(l, m[0], text+m[0], -1))

	}, nil)
	return outputs
}

// Absent remove pattern if found.
func Absent(lines []string, re *regexp.Regexp) []string {
	var outputs []string
	outputs = handleLines(lines, re, func(l string, m []string, o []string, re *regexp.Regexp) []string {
		return o
	}, nil)
	return outputs
}

// Present make sure pattern exists.
func Present(lines []string, re *regexp.Regexp, pattern string, text string) []string {
	var outputs []string
	outputs = handleLines(lines, re, nil, func(o []string, re *regexp.Regexp) []string {
		if text == "" {
			return append(o, pattern)
		}
		return append(o, text)
	})
	return outputs
}

func handleLines(lines []string, re *regexp.Regexp, matchFunc func(string, []string, []string, *regexp.Regexp) []string,
	noMatchFunc func([]string, *regexp.Regexp) []string) []string {
	var outputs []string
	matched := false
	for _, l := range lines {
		results := re.FindAllString(l, -1)
		if results != nil {
			if matchFunc != nil {
				outputs = matchFunc(l, results, outputs, re)
			} else {
				outputs = append(outputs, l)
			}
			matched = true
		} else {
			outputs = append(outputs, l)
		}
	}
	if !matched && noMatchFunc != nil {
		outputs = noMatchFunc(outputs, re)
	}
	return outputs
}
