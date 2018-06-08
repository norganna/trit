package trit

// And calculates A ∧ B ( A AND B ) for 2 trils.
func And(a, b Trilean) Trilean {
	return Min(a, b)
}

// Or calculates A ∨ B ( A OR B ) for 2 trils.
func Or(a, b Trilean) Trilean {
	return Max(a, b)
}

// Not calculates the complement ¬A ( NOT A ) for a tril.
func Not(a Trilean) Trilean {
	return Negate(a)
}

// ImpK returns the material implication of A → B ( if A then B ) for 2 trills (using Kleene logic).
func ImpK(a, b Trilean) Trilean {
	return Or(-a, b)
}

// ImpŁ returns the material implication of A → B ( if A then B ) for 2 trills (using Łukasiewicz logic).
func ImpŁ(a, b Trilean) Trilean {
	if b == Neg && a == Pos {
		return Neg
	}
	if b == Nor && a == Pos || b == Neg && a == Nor {
		return Nor
	}
	return Pos
}

// Equivalent returns the equivalence of A ↔ B ( A is equivalent to B ) for 2 trills.
func Equivalent(a, b Trilean) Trilean {
	return a * b
}

// Min returns the minimum Trilean value from the supplied parameters.
func Min(vv ...Trilean) (min Trilean) {
	min = Pos
	for _, v := range vv {
		if v < min {
			min = v
		}
	}
	return
}

// Max returns the maximum Trilean value from the supplied parameters.
func Max(vv ...Trilean) (max Trilean) {
	max = Neg
	for _, v := range vv {
		if v > max {
			max = v
		}
	}
	return
}

// Negate returns the inverse value of a tril.
func Negate(a Trilean) Trilean {
	return -a
}
