//go:generate sumgen "Expr = KeyValueExpr | *CallExpr"
package example

type Expr interface {
	expr()
}

type KeyValueExpr struct {
	Key   Expr
	Colon int
	Value Expr
}

type CallExpr struct {
	Fun      Expr
	Lparen   int
	Args     []Expr
	Ellipsis bool
	Rparen   int
}
