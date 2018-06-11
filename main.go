package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"io/ioutil"
	"os"

	"github.com/kjelly/lineinfile/exit"
	"github.com/kjelly/lineinfile/lib"
	"regexp"
	"strings"
)

func readFile(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		os.Exit(exit.FAILED_TO_READ_FILE)
	}
	return string(bytes)

}

func main() {
	var args struct {
		Path    string `arg:"positional"`
		Pattern string `arg:"positional"`
		Mode    string `arg:"-m"`
		Text    string `arg:"-t"`
	}
	args.Text = ""
	arg.MustParse(&args)

	lines := strings.Split(readFile(args.Path), "\n")

	re, err := regexp.Compile(args.Pattern)

	if err != nil {
		os.Exit(exit.INVAILED_PATTERN)
	}

	var outputs []string

	switch args.Mode {
	case "present":
		outputs = lib.Present(lines, re, args.Pattern, args.Text)
	case "absent":
		outputs = lib.Absent(lines, re)
	case "insertAfter":
		outputs = lib.InsertAfter(lines, re, args.Text)
	case "insertBefore":
		outputs = lib.InsertBefore(lines, re, args.Text)
	case "replace":
		outputs = lib.Replace(lines, re, args.Text)
	default:
		outputs = lib.Present(lines, re, args.Pattern, args.Text)

	}

	for _, l := range outputs {
		fmt.Printf("%s\n", l)
	}

}
