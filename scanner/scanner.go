package scanner

import (
	"unicode/utf8"

	"github.com/royvandewater/glox/token"
)

func New(source string) *Scanner {
	return &Scanner{
		source:  source,
		tokens:  make([]*token.Token, 0),
		start:   0,
		current: 0,
		line:    1,
	}
}

type Scanner struct {
	source  string
	tokens  []*token.Token
	start   int
	current int
	line    int
}

func (s *Scanner) ScanTokens() []*token.Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.addToken(token.EOF)

	return s.tokens
}

func (s *Scanner) scanToken() {
	runeValue := s.advance()

	switch runeValue {
	case '(':
		s.addToken(token.LEFT_PAREN)
	case ')':
		s.addToken(token.RIGHT_PAREN)
	case '{':
		s.addToken(token.LEFT_BRACE)
	case '}':
		s.addToken(token.RIGHT_BRACE)
	case ',':
		s.addToken(token.COMMA)
	case '.':
		s.addToken(token.DOT)
	case '-':
		s.addToken(token.MINUS)
	case '+':
		s.addToken(token.PLUS)
	case ';':
		s.addToken(token.SEMICOLON)
	case '*':
		s.addToken(token.STAR)
	}
}

func (s *Scanner) addToken(token_type string) {
	s.addTokenLiteral(token_type, nil)
}

func (s *Scanner) addTokenLiteral(token_type string, literal interface{}) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, token.New(token_type, text, literal, s.line))
}

func (s *Scanner) advance() rune {
	value, width := utf8.DecodeRuneInString(s.source[s.current:])

	s.current = s.current + width
	return value
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
