package token

import "fmt"

type Token struct {
	Type       string
	Lexeme     string
	Literal    interface{}
	LineNumber int
}

func New(tokenType string, lexeme string, literal interface{}, lineNumber int) *Token {
	return &Token{
		Type:       tokenType,
		Lexeme:     lexeme,
		Literal:    literal,
		LineNumber: lineNumber,
	}
}

func (t *Token) String() string {
	return fmt.Sprint(t.Type, " ", t.Lexeme, " ", t.Literal)
}
