package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var rules []string = []string{
	"Binary   : Expr left, Token operator, Expr right",
	"Grouping : Expr expression",
	"Literal  : any value",
	"Unary    : Token operator, Expr right",
}

var titleCaser = cases.Title(language.Und)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run generateast/generateast.go <output directory>")
		os.Exit(64)
	}

	outputDir := os.Args[1]
	astStr := defineAst("Expr", rules)

	outputFilePath := filepath.Join(outputDir, "expr.go")
	err := os.WriteFile(outputFilePath, []byte(astStr), 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to write to ", outputFilePath)
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}

func defineAst(baseName string, rules []string) string {
	var data strings.Builder

	fmt.Fprintf(&data, `package %v

// IMPORTANT: do not update this file directly. It is generated
// by running "go generate". The generator code can be found in
// generateast/generateast.go

import (
	token "github.com/royvandewater/glox/token"
)

type Token = token.Token

type Expr interface {
	Accept()
}

type Visitor interface {
`, strings.ToLower(baseName))

	for _, rule := range rules {
		className, _, _ := strings.Cut(rule, ":")
		className = strings.TrimSpace(className)
		fmt.Fprintf(&data, "  Visit%v%v(%v %v)\n", titleCaser.String(className), titleCaser.String(baseName), strings.ToLower(baseName), titleCaser.String(className))
	}
	fmt.Fprintln(&data, "}")
	fmt.Fprintln(&data)

	for _, rule := range rules {
		className, fieldsStr, _ := strings.Cut(rule, ":")
		className = strings.TrimSpace(className)
		fields := strings.Split(strings.TrimSpace(fieldsStr), ", ")

		fmt.Fprintln(&data, defineType(baseName, className, fields))
	}

	return data.String()
}

func defineType(baseName, className string, fields []string) string {
	var data strings.Builder

	// struct
	fmt.Fprintf(&data, "type %v struct {\n", className)
	for _, field := range fields {
		typeName, name, _ := strings.Cut(field, " ")
		fmt.Fprintf(&data, "  %v %v;\n", titleCaser.String(name), typeName)
	}
	fmt.Fprintln(&data, "}")
	fmt.Fprintln(&data)

	// constructor
	fmt.Fprintf(&data, `
func New%v(%v) *%v {
	return &%v{
`, className, formatConstructorParams(fields), className, className)

	for _, field := range fields {
		_, name, _ := strings.Cut(field, " ")
		fmt.Fprintf(&data, "    %v: %v,\n", titleCaser.String(name), name)
	}
	fmt.Fprintln(&data, "  }")
	fmt.Fprintln(&data, "}")

	// Visitor pattern
	fmt.Fprintf(&data, `
func (e *%v) Accept(visitor Visitor) {
	visitor.Visit%s%s(e)
}
`, className, titleCaser.String(className), titleCaser.String(baseName))

	return data.String()
}

func formatConstructorParams(fields []string) string {
	params := make([]string, len(fields))

	for i, field := range fields {
		className, name, _ := strings.Cut(field, " ")
		params[i] = fmt.Sprintf("%v %v", name, className)
	}

	return strings.Join(params, ", ")
}
