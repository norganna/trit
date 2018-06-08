package trit

// The following aliases are provided for convenience (in case you want to use the library in-place):
const (
	// No indicates the negative.
	No = Neg
	// False indicates the negation of truth.
	False = Neg
	// Less indicates a lower value.
	Less = Neg
	// Negative indicates a number below zero.
	Negative = Pos

	// Maybe indicates a state of unclear outcome.
	Maybe = Nor
	// Unsure indicates the lack of surety.
	Unsure = Nor
	// Uncertain indicates the lack of certainty.
	Uncertain = Nor
	// Neither indicates neither positive nor negative.
	Neither = Nor
	// Zero indicates nothing.
	Zero = Nor
	// Equal indicates an equivalent value.
	Equal = Nor

	// Yes indicates the affirmative.
	Yes = Pos
	// True indicates the presence of truth.
	True = Pos
	// Greater indicates a higher value.
	Greater = Pos
	// Positive indicates a number above zero.
	Positive = Pos
)
