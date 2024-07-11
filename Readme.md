# Enum
A library for making enums in golang using generics.

# Dependencies
- go version 1.18 and above

# Installation
run `go get github.com/eltNEG/enum`

# Usage

```go
package main

import (
	"fmt"

	"github.com/eltNEG/enum"
)

func main() {
	type dicevalue uint
	dicevalues := enum.MustNew[dicevalue](struct {
		ONE   dicevalue
		TWO   dicevalue
		THREE dicevalue
		FOUR  dicevalue
		FIVE  dicevalue
		SIX   dicevalue
	}{
		ONE:   1,
		TWO:   2,
		THREE: 3,
		FOUR:  4,
		FIVE:  5,
		SIX:   6,
	})

	// Use the dice value
	fmt.Println(dicevalues.V.FOUR) // 4

	// Verify dice value
	fmt.Println(dicevalues.IsValidValue(dicevalues.V.ONE)) // true
	fmt.Println(dicevalues.IsValidValue(dicevalue(7)))     // false

	// verify key string
	fmt.Println(dicevalues.IsValidStringKey("THREE")) // true
	fmt.Println(dicevalues.IsValidStringKey("TEN"))   // false

	// Get enum value with string
	fmt.Println(dicevalues.GetValueWithStringKey("FOUR")) // 4, true

	// Get enum value with string or panic
	fmt.Println(dicevalues.MustGetValueWithStringKey("FOUR")) // 4

	// Get enum key with value
	fmt.Println(dicevalues.GetKeyWithValue(5)) // FIVE, true

	// Get enum value with string or panic
	fmt.Println(dicevalues.MustGetKeyWithValue(6)) // SIX

}
```

# License
- MIT license.
