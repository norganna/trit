# Trit (trinary digit)

Trit is a tiny library that can add trinary values which can represent `yes`, `no` and `maybe` (or `greater`, `less` and `equal`, etc).

A Trinary digit is analogous to a binary digit (0 or 1) in that it has 3 states instead of 2 (-1, 0 or 1).

We can thus use trinary digits in the same way that binary digits are used to represent the boolean values of true and false to represent a third (middle) state of uncertainty.

To more easily use in your own code, add the following to a file called trit.go in your directory:

(Note, a more detailed version of the following is provided in `contrib/` for you to copy to your project.)

```golang
package main

import 	"github.com/nedscode/transit/lib/trit"

// Tri is an alias to the trit package's Tri value.
type Tri = trit.Tri

// The following definitions are up to your own interpretation:
const (
	yes = trit.Pos
	no = trit.Neg
	maybe = trit.Nor
)
```

## Example code

Assuming you've copied the trit.go file into your project:

```go
func canEat(thing string) Tri {
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
