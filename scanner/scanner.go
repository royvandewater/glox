package scanner

import (
	"go/token"
	"unicode/utf8"
)

func New(source string) *Scanner {
	return &Scanner{
		source:  source,
		tokens:  make([]token.Token, 0),
		start:   0,
		current: 0,
		line:    1,
	}
}

type Scanner struct {
	source  string
	tokens  []token.Token
	start   int
	current int
	line    int
}

func (s *Scanner) ScanTokens() []token.Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, token.EOF)

	return s.tokens
}

func (s *Scanner) scanToken() {

}

func (s *Scanner) advance() rune {
	value, width := utf8.DecodeRuneInString(s.source[s.current:])

	s.current = s.current + width
	return value
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
