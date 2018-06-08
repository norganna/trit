// Copy this file to your project and rename the package name below.
package YOURPACKAGE

import "github.com/norganna/trit"

// Tryte (trinary byte) is an alias to the trit package's Tryte value
type Tryte = trit.Tryte

const (
	yes   = trit.Yes
	no    = trit.No
	maybe = trit.Maybe

	greater = trit.Greater
	less    = trit.Less
	equal   = trit.Equal
)
