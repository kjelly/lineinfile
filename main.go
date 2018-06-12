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
		Path          string `arg:"positional"`
		Pattern       string `arg:"positional"`
		Mode          string `arg:"-m"`
		Text          string `arg:"-t"`
		StartLine     int
		EndLine       int
		BeforePattern string
		AfterPattern  string
		Inplace       bool
	}
	args.Text = ""
	args.StartLine = -1
	args.EndLine = -1
	args.BeforePattern = ""
	args.AfterPattern = ""
	args.Inplace = false
	arg.MustParse(&args)

	lines := strings.Split(readFile(args.Path), "\n")

	re, err := regexp.Compile(args.Pattern)

	if err != nil {
		os.Exit(exit.INVAILED_PATTERN)
	}

	var outputs []string

	p, err := lib.InitTextProcessor(args.StartLine, args.EndLine, args.BeforePattern, args.AfterPattern)

	if err != nil {
		panic(err)
	}
	switch args.Mode {
	case "present":
		outputs = p.Present(lines, re, args.Pattern, args.Text)
	case "absent":
		outputs = p.Absent(lines, re)
		outputs = p.InsertAfter(lines, re, args.Text)
		outputs = p.InsertBefore(lines, re, args.Text)
	case "replace":
		outputs = p.Replace(lines, re, args.Text)
	default:

	}

	if args.Inplace {
		ioutil.WriteFile(args.Path, []byte(strings.Join(outputs, "\n")), 0666)
	} else {
		fmt.Printf("%s", strings.Join(outputs, "\n"))
	}

}
