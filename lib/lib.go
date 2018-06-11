package lib

import (
	_ "fmt"
	"regexp"
	"strings"
)

func InitTextProcessor(startLine int, endLine int, beforePattern string,
	afterPattern string) (*TextProcessor, error) {
	re1, err := regexp.Compile(beforePattern)
	if err != nil {
		return nil, err
	}
	re2, err := regexp.Compile(afterPattern)
	if err != nil {
		return nil, err
	}
	t := &TextProcessor{
		StartLine:          startLine,
		EndLine:            endLine,
		BeforePattern:      re1,
		AfterPattern:       re2,
		inSearchArea:       false,
		lastSearchAreaLine: -1,
	}
	return t, nil

}

type TextProcessor struct {
	StartLine          int
	EndLine            int
	BeforePattern      *regexp.Regexp
	AfterPattern       *regexp.Regexp
	inSearchArea       bool
	lastSearchAreaLine int
}

// Replace replace pattern with text
func (t TextProcessor) Replace(lines []string, re *regexp.Regexp, text string) []string {
	var outputs []string
	outputs = t.handleLines(lines, re, func(l string, m []string, o []string, re *regexp.Regexp) []string {
		return append(o, strings.Replace(l, m[0], text, -1))

	}, nil)
	return outputs
}

// InsertAfter insert text after pattern.
func (t TextProcessor) InsertAfter(lines []string, re *regexp.Regexp, text string) []string {
	var outputs []string
	outputs = t.handleLines(lines, re, func(l string, m []string, o []string, re *regexp.Regexp) []string {
		return append(o, strings.Replace(l, m[0], m[0]+text, -1))

	}, nil)
	return outputs
}

// InsertBefore insert text before pattern.
func (t TextProcessor) InsertBefore(lines []string, re *regexp.Regexp, text string) []string {
	var outputs []string
	outputs = t.handleLines(lines, re, func(l string, m []string, o []string, re *regexp.Regexp) []string {
		return append(o, strings.Replace(l, m[0], text+m[0], -1))

	}, nil)
	return outputs
}

// Absent remove pattern if found.
func (t TextProcessor) Absent(lines []string, re *regexp.Regexp) []string {
	var outputs []string
	outputs = t.handleLines(lines, re, func(l string, m []string, o []string, re *regexp.Regexp) []string {
		return o
	}, nil)
	return outputs
}

// Present make sure pattern exists.
func (t TextProcessor) Present(lines []string, re *regexp.Regexp, pattern string, text string) []string {
	var outputs []string
	outputs = t.handleLines(lines, re, nil, func(o []string, re *regexp.Regexp, p int) []string {
		if text == "" {
			return insertIntoLines(p, pattern, o)
		}
		return insertIntoLines(p, text, o)
	})
	return outputs
}

func (t *TextProcessor) skip(i int, line string) bool {
	if t.StartLine != -1 && i < t.StartLine {
		return true
	}
	if t.EndLine != -1 && i > t.StartLine {
		return true
	}
	results := t.AfterPattern.FindAllString(line, -1)
	if results != nil {
		t.inSearchArea = true
	}
	results = t.BeforePattern.FindAllString(line, -1)
	if t.inSearchArea && results != nil {
		t.inSearchArea = false
		t.lastSearchAreaLine = i
	}
	return t.inSearchArea
}

func insertIntoLines(index int, line string, lines []string) []string {
	var outputs []string
	for i, l := range lines {
		if index == i {
			outputs = append(outputs, line)
		}
		outputs = append(outputs, l)
	}
	return outputs

}

func maxInt(v int, arr ...int) int {
	max := v
	for _, l := range arr {
		if l > max {
			max = l
		}
	}
	return max
}

func (t TextProcessor) handleLines(lines []string, re *regexp.Regexp, matchFunc func(string, []string, []string, *regexp.Regexp) []string,
	noMatchFunc func([]string, *regexp.Regexp, int) []string) []string {
	var outputs []string
	matched := false
	for i, l := range lines {

		if t.skip(i, l) {
			outputs = append(outputs, l)
		} else {
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
	}
	if !matched && noMatchFunc != nil {
		p := maxInt(t.EndLine-1, t.lastSearchAreaLine)
		if p == -1 {
			p = len(outputs) - 1
		}
		outputs = noMatchFunc(outputs, re, p)
	}
	return outputs
}
