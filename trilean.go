package trit

// Trilean is a trivalent value that represents yes, no or maybe / greater, less or equal / anything else you want.
type Trilean int8

const (
	// Neg is no, negative or -1.
	Neg Trilean = iota - 1
	// Nor is neither positive, nor negative, nothing, zero, maybe, unsure or 0.
	Nor
	// Pos is yes, positive or 1.
	Pos
)

// And is a convenience method to call the And(a,b) operator.
func (a Trilean) And(b Trilean) Trilean {
	return And(a, b)
}

// Or is a convenience method to call the Or(a,b) operator.
func (a Trilean) Or(b Trilean) Trilean {
	return Or(a, b)
}

// ImpK is a convenience method to call the ImpK(a,b) operator.
func (a Trilean) ImpK(b Trilean) Trilean {
	return ImpK(a, b)
}

// ImpŁ is a convenience method to call the ImpŁ(a,b) operator.
func (a Trilean) ImpŁ(b Trilean) Trilean {
	return ImpŁ(a, b)
}

// Equivalent is a convenience method to call the Equivalent(a,b) operator.
func (a Trilean) Equivalent(b Trilean) Trilean {
	return Equivalent(a, b)
}

// Not is a convenience method to call the Not(a) operator.
func (a Trilean) Not() Trilean {
	return Negate(a)
}
