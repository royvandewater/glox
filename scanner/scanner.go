package scanner

import (
	"fmt"
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

func (s *Scanner) ScanTokens() ([]*token.Token, []error) {
	errs := make([]error, 0)

	for !s.isAtEnd() {
		s.start = s.current
		err := s.scanToken()
		if err != nil {
			errs = append(errs, err)
		}

	}

	s.addToken(token.EOF)

	return s.tokens, errs
}

func (s *Scanner) scanToken() error {
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
	case '!':
		if s.match('=') {
			s.addToken(token.BANG_EQUAL)
		} else {
			s.addToken(token.BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(token.EQUAL_EQUAL)
		} else {
			s.addToken(token.EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addToken(token.LESS_EQUAL)
		} else {
			s.addToken(token.LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(token.GREATER_EQUAL)
		} else {
			s.addToken(token.GREATER)
		}

	default:
		return fmt.Errorf("unexpected character: %v", string(runeValue))
	}

	return nil
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

func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}

	index := s.current + 1
	value, _ := utf8.DecodeRuneInString(s.source[index:])

	if value != expected {
		return false
	}

	s.advance()
	return true
}
