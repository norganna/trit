package trit

/*
You can use the values directly from this library, however to more easily use in your own code, add the following
to a file called `trit.go` in your own project's directory:

```golang
import 	"github.com/nedscode/transit/lib/trit"

// Tryte is an alias to the tri package's Tryte value
type Tryte = trit.Tryte
const (
	yes = trit.Yes
	no = trit.No
	maybe = trit.Maybe

	greater = trit.Greater
	less = trit.Less
	equal = trit.Equal
)
```

*/

// Tryte is a trinary value that represents yes, no or maybe / greater, less or equal.
type Tryte int8

const (
	// No or -1
	No Tryte = iota - 1
	// Maybe or 0
	Maybe
	// Yes or 1
	Yes
)

const (
	// Less or -1
	Less Tryte = iota - 1
	// Equal or 0
	Equal
	// Greater or 1
	Greater
)
