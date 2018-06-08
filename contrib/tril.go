// Copy this file to your project and rename the package name below.
package YOURPACKAGE

import "github.com/norganna/trit"

// A tril (analogous as bool is to Boolean) is an alias to the trit package's Trilean value.
type tril = trit.Trilean

// Keep whichever of the following you'd prefer / want to use:
const (
	yes   = trit.Pos
	no    = trit.Neg
	maybe = trit.Nor
)

// For comparisons:
const (
	greater = trit.Pos
	less    = trit.Neg
	equal   = trit.Nor
)

// When dealing with veracity:
const (
	truth     = trit.Pos
	fallicy   = trit.Neg
	uncertain = trit.Nor
)

// Or make up your own:
const (
	giraffe  = trit.Pos
	elephant = trit.Neg
	mouse    = trit.Nor
)
