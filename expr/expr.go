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

type Visitor interface {
	VisitBinaryExpr(expr Binary)
	VisitGroupingExpr(expr Grouping)
	VisitLiteralExpr(expr Literal)
	VisitUnaryExpr(expr Unary)
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

func (e *Binary) Accept(visitor Visitor) {
	visitor.VisitBinaryExpr(e)
}

type Grouping struct {
	Expression Expr
}

func NewGrouping(expression Expr) *Grouping {
	return &Grouping{
		Expression: expression,
	}
}

func (e *Grouping) Accept(visitor Visitor) {
	visitor.VisitGroupingExpr(e)
}

type Literal struct {
	Value any
}

func NewLiteral(value any) *Literal {
	return &Literal{
		Value: value,
	}
}

func (e *Literal) Accept(visitor Visitor) {
	visitor.VisitLiteralExpr(e)
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

func (e *Unary) Accept(visitor Visitor) {
	visitor.VisitUnaryExpr(e)
}
