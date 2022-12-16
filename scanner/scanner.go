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
	case '/':
		if s.match('/') {
			for !s.isAtEnd() && s.peek() != '\n' {
				s.advance()
			}
		} else {
			s.addToken(token.SLASH)
		}
	case ' ': // Go case statements **DO NOT** fall through
	case '\r':
	case '\t':
	case '\n':
		s.line += 1
	case '"':
		s.parseString()

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

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return '\x00'
	}

	value, _ := utf8.DecodeRuneInString(s.source[s.current:])
	return value
}

func (s *Scanner) parseString() error {
	for !s.isAtEnd() && s.peek() != '"' {
		if s.peek() == '\n' {
			s.line += 1
		}

		s.advance()
	}

	if s.isAtEnd() {
		return fmt.Errorf("unterminated string")
	}

	// the closing "
	s.advance()

	// Trim the surrounding quotes
	value := s.source[s.start+1 : s.current-1]
	s.addTokenLiteral(token.STRING, value)
	return nil
}
