package parser

import (
	"fmt"

	"github.com/royvandewater/glox/expr"
	"github.com/royvandewater/glox/token"
)

type Parser struct {
	current int
	tokens  []*token.Token
}

func New(tokens []*token.Token) *Parser {
	return &Parser{
		current: 0,
		tokens:  tokens,
	}
}

func (p *Parser) Parse() (expr.Expr, error) {
	return p.expression()
}

func (p *Parser) expression() (expr.Expr, error) {
	return p.equality()
}

func (p *Parser) equality() (expr.Expr, error) {
	expression, err := p.comparison()
	if err != nil {
		return nil, err
	}

	for p.check(token.BANG_EQUAL, token.EQUAL_EQUAL) {
		operator := p.advance()
		right, err := p.comparison()
		if err != nil {
			return nil, err
		}

		expression = expr.NewBinary(expression, operator, right)
	}

	return expression, nil
}

func (p *Parser) comparison() (expr.Expr, error) {
	expression, err := p.term()
	if err != nil {
		return nil, err
	}

	for p.check(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.advance()
		right, err := p.term()
		if err != nil {
			return nil, err
		}

		expression = expr.NewBinary(expression, operator, right)
	}

	return expression, nil
}

func (p *Parser) term() (expr.Expr, error) {
	expression, err := p.factor()
	if err != nil {
		return nil, err
	}

	for p.check(token.PLUS, token.MINUS) {
		operator := p.advance()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}

		expression = expr.NewBinary(expression, operator, right)
	}

	return expression, nil
}

func (p *Parser) factor() (expr.Expr, error) {
	expression, err := p.unary()
	if err != nil {
		return nil, err
	}

	for p.check(token.SLASH, token.STAR) {
		operator := p.advance()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}

		expression = expr.NewBinary(expression, operator, right)
	}

	return expression, nil
}

func (p *Parser) unary() (expr.Expr, error) {
	if !p.check(token.BANG, token.MINUS) {
		return p.primary()
	}

	operator := p.advance()
	right, err := p.unary()
	if err != nil {
		return nil, err
	}

	return expr.NewUnary(operator, right), nil
}

func (p *Parser) primary() (expr.Expr, error) {
	if p.check(token.FALSE) {
		p.advance()
		return expr.NewLiteral(false), nil
	}

	if p.check(token.TRUE) {
		p.advance()
		return expr.NewLiteral(true), nil
	}

	if p.check(token.NIL) {
		p.advance()
		return expr.NewLiteral(nil), nil
	}

	if p.check(token.NUMBER, token.STRING) {
		return expr.NewLiteral(p.advance().Literal), nil
	}

	if p.check(token.LEFT_PAREN) {
		p.advance()
		expression, err := p.expression()
		if err != nil {
			return nil, err
		}

		err = p.consume(token.RIGHT_PAREN, "Expect ')' after expression.")
		if err != nil {
			return nil, err
		}

		return expr.NewGrouping(expression), nil
	}

	return nil, fmt.Errorf("unrecognized token: %v", p.peek())
}

func (p *Parser) advance() token.Token {
	val := p.peek()
	p.current += 1

	return val
}

func (p *Parser) check(tokenTypes ...string) bool {
	if p.isAtEnd() {
		return false
	}

	nextToken := p.peek()

	for _, tokenType := range tokenTypes {
		if nextToken.Type == tokenType {
			return true
		}
	}

	return false
}

func (p *Parser) consume(tokenType, message string) error {
	if p.check(tokenType) {
		p.advance()
		return nil
	}

	return fmt.Errorf("could not consume: %v. \"%v\"", p.peek(), message)
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == token.EOF
}

func (p *Parser) peek() token.Token {
	return *p.tokens[p.current]
}
