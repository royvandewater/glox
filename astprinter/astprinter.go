package astprinter

import (
	"fmt"
	"strings"

	"github.com/royvandewater/glox/expr"
)

type AstPrinter struct {
	ReturnValue string
}

func (a *AstPrinter) Print(expr expr.Expr) string {
	expr.Accept(a)
	return a.ReturnValue
}

// VisitBinaryExpr implements expr.Visitor
func (a *AstPrinter) VisitBinaryExpr(expr *expr.Binary) {
	a.ReturnValue = a.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

// VisitGroupingExpr implements expr.Visitor
func (a *AstPrinter) VisitGroupingExpr(expr *expr.Grouping) {
	a.ReturnValue = a.parenthesize("group", expr.Expression)
}

// VisitLiteralExpr implements expr.Visitor
func (a *AstPrinter) VisitLiteralExpr(expr *expr.Literal) {
	if expr.Value == nil {
		a.ReturnValue = "nil"
		return
	}

	a.ReturnValue = fmt.Sprintf("%v", expr.Value)
}

// VisitUnaryExpr implements expr.Visitor
func (a *AstPrinter) VisitUnaryExpr(expr *expr.Unary) {
	a.ReturnValue = a.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (a *AstPrinter) parenthesize(name string, exprs ...expr.Expr) string {
	var data strings.Builder

	fmt.Fprint(&data, "(", name)

	for _, expr := range exprs {
		fmt.Fprint(&data, " ", a.Print(expr))
	}

	fmt.Fprint(&data, ")")

	return data.String()
}
