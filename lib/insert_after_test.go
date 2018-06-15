package lib

import (
	"strings"
	"testing"
)

func TestTextInsertAfter(t *testing.T) {
	pattern := "user"
	p, err := InitTextProcessor(pattern, "", "name", -1, -1, "", "")

	if err != nil {
		t.Errorf("%s", err)
	}

	o := strings.Join(p.InsertAfter(strings.Split(sample, "\n")), "\n")

	output := `
[Default]
username=abc
password=abc
url=http://www.example.com
[auth]
username=ccc
password=ccc

[data]
url=http://example.com
uri=http://www.example.com
driver=sql
`
	if o != output {
		t.Errorf("Error %s, %s\n", o, output)
	}

}

func TestTextInsertAfterBetweenPattern(t *testing.T) {
	pattern := "user"
	p, err := InitTextProcessor(pattern, "", "name", -1, -1, "\\[auth\\]", "\\[data\\]")

	if err != nil {
		t.Errorf("%s", err)
	}

	o := strings.Join(p.InsertAfter(strings.Split(sample, "\n")), "\n")

	output := `
[Default]
user=abc
password=abc
url=http://www.example.com
[auth]
username=ccc
password=ccc

[data]
url=http://example.com
uri=http://www.example.com
driver=sql
`
	if o != output {
		t.Errorf("Error %s, %s\n", o, output)
	}

}

func TestTextInsertAfterBetweenLineNumber(t *testing.T) {
	pattern := "user"
	p, err := InitTextProcessor(pattern, "", "name", 4, 7, "", "")

	if err != nil {
		t.Errorf("%s", err)
	}

	o := strings.Join(p.InsertAfter(strings.Split(sample, "\n")), "\n")

	output := `
[Default]
user=abc
password=abc
url=http://www.example.com
[auth]
username=ccc
password=ccc

[data]
url=http://example.com
uri=http://www.example.com
driver=sql
`
	if o != output {
		t.Errorf("Error %s, %s\n", o, output)
	}

}

func TestTextInsertAfterPatternNotFoundBetweenLineNumber(t *testing.T) {
	pattern := "drive"
	p, err := InitTextProcessor(pattern, "", "name", 4, 7, "", "")

	if err != nil {
		t.Errorf("%s", err)
	}

	o := strings.Join(p.InsertAfter(strings.Split(sample, "\n")), "\n")

	output := `
[Default]
user=abc
password=abc
url=http://www.example.com
[auth]
user=ccc
password=ccc

[data]
url=http://example.com
uri=http://www.example.com
driver=sql
`
	if o != output {
		t.Errorf("Error %s, %s\n", o, output)
	}

}

func TestTextInsertAfterPatternNotFoundBetweenPattern(t *testing.T) {
	pattern := "drive"
	p, err := InitTextProcessor(pattern, "", "name", -1, -1, "\\[auth\\]", "\\[data\\]")

	if err != nil {
		t.Errorf("%s", err)
	}

	o := strings.Join(p.InsertAfter(strings.Split(sample, "\n")), "\n")

	output := `
[Default]
user=abc
password=abc
url=http://www.example.com
[auth]
user=ccc
password=ccc

[data]
url=http://example.com
uri=http://www.example.com
driver=sql
`
	if o != output {
		t.Errorf("Error %s, %s\n", o, output)
	}

}

func TestTextInsertAfterPatternNotFoundBeforeLineNumber(t *testing.T) {
	pattern := "drive"
	p, err := InitTextProcessor(pattern, "", "name", -1, 10, "", "")

	if err != nil {
		t.Errorf("%s", err)
	}

	o := strings.Join(p.InsertAfter(strings.Split(sample, "\n")), "\n")

	output := `
[Default]
user=abc
password=abc
url=http://www.example.com
[auth]
user=ccc
password=ccc

[data]
url=http://example.com
uri=http://www.example.com
driver=sql
`
	if o != output {
		t.Errorf("Error %s, %s\n", o, output)
	}

}
