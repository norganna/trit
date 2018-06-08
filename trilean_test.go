package trit_test

import (
	"testing"

	"github.com/norganna/trit"
)

type tril = trit.Trilean

const (
	yes   = trit.Pos
	no    = trit.Neg
	maybe = trit.Nor
)

func unary(t *testing.T, name string, op func(tril, tril) tril, a, b, expect tril) {
	got := op(a, b)
	if got != expect {
		t.Errorf(
			"Error performing %s with %d, %d. Got %d, Expected %d",
			name,
			a,
			b,
			got,
			expect,
		)
	}
}

func TestTrilean_Operators(t *testing.T) {
	unary(t, "and", trit.And, yes, yes, yes)
	unary(t, "and", trit.And, yes, no, no)
	unary(t, "and", trit.And, no, yes, no)
	unary(t, "and", trit.And, yes, maybe, maybe)
	unary(t, "and", trit.And, maybe, yes, maybe)
	unary(t, "and", trit.And, maybe, maybe, maybe)
	unary(t, "and", trit.And, no, maybe, no)
	unary(t, "and", trit.And, maybe, no, no)
	unary(t, "and", trit.And, no, no, no)

	unary(t, "or", trit.Or, yes, yes, yes)
	unary(t, "or", trit.Or, yes, no, yes)
	unary(t, "or", trit.Or, no, yes, yes)
	unary(t, "or", trit.Or, yes, maybe, yes)
	unary(t, "or", trit.Or, maybe, yes, yes)
	unary(t, "or", trit.Or, maybe, maybe, maybe)
	unary(t, "or", trit.Or, no, maybe, maybe)
	unary(t, "or", trit.Or, maybe, no, maybe)
	unary(t, "or", trit.Or, no, no, no)

	unary(t, "impK", trit.ImpK, yes, yes, yes)
	unary(t, "impK", trit.ImpK, yes, no, no)
	unary(t, "impK", trit.ImpK, no, yes, yes)
	unary(t, "impK", trit.ImpK, yes, maybe, maybe)
	unary(t, "impK", trit.ImpK, maybe, yes, yes)
	unary(t, "impK", trit.ImpK, maybe, maybe, maybe)
	unary(t, "impK", trit.ImpK, no, maybe, yes)
	unary(t, "impK", trit.ImpK, maybe, no, maybe)
	unary(t, "impK", trit.ImpK, no, no, yes)

	unary(t, "impŁ", trit.ImpŁ, yes, yes, yes)
	unary(t, "impŁ", trit.ImpŁ, yes, no, no)
	unary(t, "impŁ", trit.ImpŁ, no, yes, yes)
	unary(t, "impŁ", trit.ImpŁ, yes, maybe, maybe)
	unary(t, "impŁ", trit.ImpŁ, maybe, yes, yes)
	unary(t, "impŁ", trit.ImpŁ, maybe, maybe, yes)
	unary(t, "impŁ", trit.ImpŁ, no, maybe, yes)
	unary(t, "impŁ", trit.ImpŁ, maybe, no, maybe)
	unary(t, "impŁ", trit.ImpŁ, no, no, yes)

	unary(t, "equiv", trit.Equivalent, yes, yes, yes)
	unary(t, "equiv", trit.Equivalent, yes, no, no)
	unary(t, "equiv", trit.Equivalent, no, yes, no)
	unary(t, "equiv", trit.Equivalent, yes, maybe, maybe)
	unary(t, "equiv", trit.Equivalent, maybe, yes, maybe)
	unary(t, "equiv", trit.Equivalent, maybe, maybe, maybe)
	unary(t, "equiv", trit.Equivalent, no, maybe, maybe)
	unary(t, "equiv", trit.Equivalent, maybe, no, maybe)
	unary(t, "equiv", trit.Equivalent, no, no, yes)
}
