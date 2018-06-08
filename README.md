# Trit (trinary digit)

Trit is a tiny library that can add trinary values which can represent `yes`, `no` and `maybe` (or `greater`, `less` and `equal`, etc).

A Trinary digit is analogous to a binary digit (0 or 1) in that it has 3 states instead of 2 (-1, 0 or 1).

We can thus use trinary digits in the same way that binary digits are used to represent the boolean values of true and false to represent a third (middle) state of uncertainty.

To more easily use in your own code, add the file as provided in `contrib/tril.go` in your directory.

## Example code

```golang
package main

import 	"github.com/nedscode/transit/lib/trit"

// A tril (analogous as bool is to Boolean) is an alias to the trit package's Trilean value.
type tril = trit.Trilean

// The following definitions are up to your own interpretation:
const (
	yes = trit.Pos
	no = trit.Neg
	maybe = trit.Nor
)

func canEat(thing string) tril {
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
