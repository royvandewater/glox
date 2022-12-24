package expr

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

type Binary struct {
	Left     Expr
	Operator Token
	Right    Expr
}

func NewBinary(left Expr, operator Token, right Expr) *Binary {
	return &Binary{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}

type Grouping struct {
	Expression Expr
}

func NewGrouping(expression Expr) *Grouping {
	return &Grouping{
		Expression: expression,
	}
}

type Literal struct {
	Value any
}

func NewLiteral(value any) *Literal {
	return &Literal{
		Value: value,
	}
}

type Unary struct {
	Operator Token
	Right    Expr
}

func NewUnary(operator Token, right Expr) *Unary {
	return &Unary{
		Operator: operator,
		Right:    right,
	}
}
