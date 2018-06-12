package lib

import (
	"regexp"
	"strings"
)

// InitTextProcessor init TextProcessor
func InitTextProcessor(pattern string, replacedText string, text string, startLine int, endLine int, afterPattern string,
	beforePattern string) (*TextProcessor, error) {
	var beforePatternRe, afterPatternRe *regexp.Regexp
	var err error
	if beforePattern == "" {
		beforePatternRe = nil
	} else {
		beforePatternRe, err = regexp.Compile(beforePattern)
		if err != nil {
			return nil, err
		}
	}
	if afterPattern == "" {
		afterPatternRe = nil
	} else {
		afterPatternRe, err = regexp.Compile(afterPattern)
		if err != nil {
			return nil, err
		}
	}
	patternRe, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	var replacedTextRe *regexp.Regexp
	if replacedText == "" {
		replacedText = pattern
		replacedTextRe = patternRe
	} else {
		replacedTextRe, err = regexp.Compile(replacedText)
		if err != nil {
			return nil, err
		}
	}
	t := &TextProcessor{
		StartLine:          startLine,
		EndLine:            endLine,
		BeforePatternRe:    beforePatternRe,
		AfterPatternRe:     afterPatternRe,
		pattern:            pattern,
		patternRe:          patternRe,
		replacedText:       replacedText,
		replacedTextRe:     replacedTextRe,
		text:               text,
		inSearchArea:       false,
		lastSearchAreaLine: -1,
	}
	t.reset()
	return t, nil

}

// TextProcessor used for processing text
type TextProcessor struct {
	StartLine          int
	EndLine            int
	BeforePatternRe    *regexp.Regexp
	AfterPatternRe     *regexp.Regexp
	pattern            string
	patternRe          *regexp.Regexp
	replacedText       string
	replacedTextRe     *regexp.Regexp
	text               string
	inSearchArea       bool
	lastSearchAreaLine int
}

func (t *TextProcessor) reset() {
	t.inSearchArea = false
	t.lastSearchAreaLine = -1
}

// Replace replace pattern with text
func (t TextProcessor) Replace(lines []string) []string {
	var outputs []string
	outputs = t.handleLines(lines, t.patternRe, func(l string, m []string, o []string) []string {
		return append(o, string(t.patternRe.ReplaceAllString(l, t.text)))

	}, nil)
	return outputs
}

// InsertAfter insert text after pattern.
func (t TextProcessor) InsertAfter(lines []string) []string {
	var outputs []string
	outputs = t.handleLines(lines, t.patternRe, func(l string, m []string, o []string) []string {
		return append(o, strings.Replace(l, m[0], m[0]+t.text, -1))
	}, nil)
	return outputs
}

// InsertBefore insert text before pattern.
func (t TextProcessor) InsertBefore(lines []string) []string {
	var outputs []string
	outputs = t.handleLines(lines, t.patternRe, func(l string, m []string, o []string) []string {
		return append(o, strings.Replace(l, m[0], t.text+m[0], -1))

	}, nil)
	return outputs
}

// Absent remove pattern if found.
func (t TextProcessor) Absent(lines []string) []string {
	var outputs []string
	outputs = t.handleLines(lines, t.patternRe, func(l string, m []string, o []string) []string {
		return o
	}, nil)
	return outputs
}

// Present make sure pattern exists.
func (t TextProcessor) Present(lines []string) []string {
	var outputs []string
	outputs = t.handleLines(lines, t.patternRe, nil, func(o []string, p int) []string {
		if t.text == "" {
			return insertIntoLines(p, t.pattern, o)
		}
		return insertIntoLines(p, t.text, o)
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
	if t.AfterPatternRe == nil && t.BeforePatternRe == nil {
		return false
	}

	var results []string
	if t.AfterPatternRe != nil {
		results = t.AfterPatternRe.FindAllString(line, -1)
		if results != nil {
			t.inSearchArea = true
			return true
		}
	}

	if t.BeforePatternRe != nil {
		results = t.BeforePatternRe.FindAllString(line, -1)
		if t.inSearchArea && results != nil {
			t.inSearchArea = false
			t.lastSearchAreaLine = i
		}

	}

	return !t.inSearchArea
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

func (t TextProcessor) handleLines(lines []string, re *regexp.Regexp, matchFunc func(string, []string, []string) []string,
	noMatchFunc func([]string, int) []string) []string {
	var outputs []string
	matched := false
	for i, l := range lines {

		if t.skip(i, l) {
			outputs = append(outputs, l)
		} else {
			results := re.FindAllString(l, -1)
			if results != nil {
				if matchFunc != nil {
					outputs = matchFunc(l, results, outputs)
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
		outputs = noMatchFunc(outputs, p)
	}
	return outputs
}
