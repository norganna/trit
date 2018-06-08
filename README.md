# Trit (trinary digit)

Trit is a tiny library that can add trinary values which can represent `yes`, `no` and `maybe` (or `greater`, `less` and `equal` if you prefer).

To more easily use in your own code, add the following to a file called trit.go in your directory:

```golang
import 	"github.com/nedscode/transit/lib/trit"

// Tryte (trinary byte) is an alias to the tri package's Tryte value
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

## Example code

Assuming you've copied the trit.go file into your project:

```go
func canEat(thing string) Tryte {
	switch thing {
    case "nail":
        return no
    case "cake":
        return yes
    default:
        return maybe
	}
}
```
