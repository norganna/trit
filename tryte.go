package trit

// Tryte is a trinary value that represents yes, no or maybe / greater, less or equal.
type Tryte int8

const (
	// Neg is negative or -1.
	Neg Tryte = iota - 1
	// Nor is neither positive, nor negative, nothing, zero or 0.
	Nor
	// Pos is positive or 1.
	Pos
)

// The following aliases are provided for convenience, if you want to use the library in-place:
const (
	// No indicates the negative.
	No = Neg
	// False indicates the negative.
	False = Neg
	// Less indicates a lower value.
	Less = Neg
	// Negative indicates a number below zero.
	Negative = Pos

	// Maybe indicates unsuredness.
	Maybe = Nor
	// Unsure indicates unsuredness.
	Unsure = Nor
	// Uncertain indicates unsuredness.
	Uncertain = Nor
	// Neither indicates neither positive nor negative.
	Neither = Nor
	// Zero indicates nothing.
	Zero = Nor
	// Equal indicates an equivalent value.
	Equal = Nor

	// Yes indicates the affirmative.
	Yes = Pos
	// True indicates the affirmative.
	True = Pos
	// Greater indicates a higher value.
	Greater = Pos
	// Positive indicates a number above zero.
	Positive = Pos
)
