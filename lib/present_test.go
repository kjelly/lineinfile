package lib

import (
	"strings"
	"testing"
)

func TestTextPresent(t *testing.T) {
	pattern := "user=abc"
	p, err := InitTextProcessor(pattern, "", "", -1, -1, "", "")

	if err != nil {
		t.Errorf("Failed to init InitTextProcessor. Reason: %s", err.Error())
	}

	o := strings.Join(p.Present(strings.Split(sample, "\n")), "\n")
	if o != sample {
		t.Errorf("Error %s, %s\n", o, sample)
	}
}

func TestTextPresentNotFound(t *testing.T) {
	pattern := "method=post"
	p, err := InitTextProcessor(pattern, "", "", -1, -1, "", "")

	if err != nil {
		t.Errorf("Failed to init InitTextProcessor. Reason: %s", err.Error())
	}

	o := strings.Join(p.Present(strings.Split(sample, "\n")), "\n")
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
method=post
`
	if o != output {
		t.Errorf("Error %s, %s\n", o, output)
	}
}

func TestTextPresentNotFoundAfterLine(t *testing.T) {
	pattern := "user=abc"
	p, err := InitTextProcessor(pattern, "", "", 3, -1, "", "")

	if err != nil {
		t.Errorf("Failed to init InitTextProcessor. Reason: %s", err.Error())
	}

	o := strings.Join(p.Present(strings.Split(sample, "\n")), "\n")
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
user=abc
`
	if o != output {
		t.Errorf("Error %s, %s\n", o, output)
	}
}

func TestTextPresentNotFounBetweenLine(t *testing.T) {
	pattern := "url=sql://example.com"
	p, err := InitTextProcessor(pattern, "", "", 1, 5, "", "")

	if err != nil {
		t.Errorf("Failed to init InitTextProcessor. Reason: %s", err.Error())
	}

	if err != nil {
		t.Errorf("%s\n", err)
	}

	o := strings.Join(p.Present(strings.Split(sample, "\n")), "\n")
	output := `
[Default]
user=abc
password=abc
url=sql://example.com
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
func TestTextPresentNotFounBetweenPattern(t *testing.T) {
	pattern := "url=sql://example.com"
	p, err := InitTextProcessor(pattern, "", "", -1, -1,
		"\\[auth\\]", "\\[data\\]")

	if err != nil {
		t.Errorf("Failed to init InitTextProcessor. Reason: %s", err.Error())
	}

	o := strings.Join(p.Present(strings.Split(sample, "\n")), "\n")
	output := `
[Default]
user=abc
password=abc
url=http://www.example.com
[auth]
user=ccc
password=ccc

url=sql://example.com
[data]
url=http://example.com
uri=http://www.example.com
driver=sql
`
	if o != output {
		t.Errorf("Error %s, %s\n", o, output)
	}
}
