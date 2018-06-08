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
