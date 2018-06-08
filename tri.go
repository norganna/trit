package trit

// Tri is a trinary value that represents yes, no or maybe / greater, less or equal.
type Tri int8

const (
	// Neg is negative or -1.
	Neg Tri = iota - 1
	// Nor is neither positive, nor negative, nothing, zero or 0.
	Nor
	// Pos is positive or 1.
	Pos
)
