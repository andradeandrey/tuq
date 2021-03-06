package parser

import (
	//"log"
	"fmt"
)

type GreaterThanOrEqualExpression struct {
	Left  Expression
	Right Expression
}

func NewGreaterThanOrEqualExpression(l, r Expression) *GreaterThanOrEqualExpression {
	return &GreaterThanOrEqualExpression{
		Left:  l,
		Right: r}
}

func (gte *GreaterThanOrEqualExpression) String() string {
	return fmt.Sprintf("%v >= %v", gte.Left, gte.Right)
}

func (gte *GreaterThanOrEqualExpression) SymbolsReferenced() []string {
	leftSymbols := gte.Left.SymbolsReferenced()
	return concatStringSlices(leftSymbols, gte.Right.SymbolsReferenced())
}

func (gte *GreaterThanOrEqualExpression) PrefixSymbols(s string) {
	gte.Left.PrefixSymbols(s)
	gte.Right.PrefixSymbols(s)
}
