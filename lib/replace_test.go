package lib

import (
	"strings"
	"testing"
)

func TestTextReplace(t *testing.T) {
	pattern := "(user.*)"
	p, _ := InitTextProcessor(pattern, "", "os_$1", -1, -1, "", "")

	o := strings.Join(p.Replace(strings.Split(sample, "\n")), "\n")

	output := `
[Default]
os_user=abc
password=abc
url=http://www.example.com
[auth]
os_user=ccc
password=ccc

[data]
url=http://example.com
uri=http://www.example.com
driver=sql
`
	if o != output {
		t.Errorf("Error %s, %s\n", o, sample)
	}

}
