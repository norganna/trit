// Copy this file to your project and rename the package name below.
package YOURPACKAGE

import "github.com/norganna/trit"

// Tryte (trinary byte) is an alias to the trit package's Tryte value.
type Tryte = trit.Tryte

// Keep whichever of the following you'd prefer / want to use:
const (
	yes   = trit.Pos
	no    = trit.Neg
	maybe = trit.Nor
)

const (
	greater = trit.Pos
	less    = trit.Neg
	equal   = trit.Nor
)

const (
	truth       = trit.Pos
	falsity     = trit.Neg
	uncertainty = trit.Nor
)

// Or make up your own:
const (
	giraffe  = trit.Pos
	elephant = trit.Neg
	mouse    = trit.Nor
)
