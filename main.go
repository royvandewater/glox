package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/royvandewater/glox/scanner"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Fprintln(os.Stderr, " usage: glox [script]")
		os.Exit(64)
	}

	if len(os.Args) == 2 {
		err := runFile(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error running file: ", err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}

	runPrompt()
	os.Exit(0)
}

func runFile(filename string) error {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return err
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

		err = run(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error running line: ", err.Error())
		}
	}
}

func run(source string) error {
	scanner := scanner.New(source)
	tokens := scanner.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}
	return nil
}
