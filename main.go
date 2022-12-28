package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/royvandewater/glox/astprinter"
	"github.com/royvandewater/glox/parser"
	"github.com/royvandewater/glox/scanner"
)

// command to regenerate ast
//go:generate sh -c "go run generateast/generateast.go expr/ && go fmt ./expr/ > /dev/null"

func main() {
	if len(os.Args) > 2 {
		fmt.Fprintln(os.Stderr, " usage: glox [script]")
		os.Exit(64)
	}

	if len(os.Args) == 2 {
		errs := runFile(os.Args[1])
		if len(errs) > 0 {
			fmt.Fprintln(os.Stderr, "Error running file: ")
			for _, err := range errs {
				fmt.Fprintln(os.Stderr, err.Error())
			}
			os.Exit(1)
		}
		os.Exit(0)
	}

	runPrompt()
	os.Exit(0)
}

func runFile(filename string) []error {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return []error{err}
	}

	return run(string(bytes))
}

func runPrompt() error {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		errs := run(line)
		for _, err := range errs {
			fmt.Fprintln(os.Stderr, "Error running line: ", err.Error())
		}
	}
}

func run(source string) []error {
	scanner := scanner.New(source)
	tokens, errs := scanner.ScanTokens()

	if len(errs) > 0 {
		return errs
	}

	parser := parser.New(tokens)
	expr, err := parser.Parse()

	if err != nil {
		return []error{err}
	}

	fmt.Println(astprinter.New().Print(expr))
	return nil
}
