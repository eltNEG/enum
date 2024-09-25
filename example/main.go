package main

import (
	"fmt"

	"github.com/eltNEG/enum"
)

func main() {
	// Simple enum
	type weekday uint8
	weekdays := enum.Make[weekday](struct {
		Monday,
		Tuesday,
		Wednesday,
		Thursday,
		Friday,
		Saturday,
		Sunday weekday
	}{})

	// Use the enum
	fmt.Println(weekdays.V().Monday) // 0

	// Verify enum value
	fmt.Println(weekdays.IsValidValue(weekdays.V().Monday)) // true

	// Verify key string
	fmt.Println(weekdays.IsValidStringKey("Tuesday")) // true

	// Get enum value with string
	fmt.Println(weekdays.GetValueWithStringKey("Friday")) // 4, true

	// Get enum key with index
	fmt.Println(weekdays.GetKeyAtIndex(5)) // Saturday, true

	// Get enum with out of range index
	fmt.Println(weekdays.GetKeyAtIndex(7)) // "", false

	// Get all the values as a slice (unordered)
	fmt.Println(weekdays.Values()) // [0 1 2 3 4 5 6]

	// Get all the keys as a slice (unordered)
	fmt.Println(weekdays.Keys()) // [Monday Tuesday Wednesday Thursday Friday Saturday Sunday]

	fmt.Println("\n\nEnum with custom value type")
	// Custom enum value type
	type dicevalue uint
	dicevalues := enum.MustNew[dicevalue](struct {
		ONE,
		TWO,
		THREE,
		FOUR,
		FIVE,
		SIX dicevalue
	}{
		ONE:   1,
		TWO:   2,
		THREE: 3,
		FOUR:  4,
		FIVE:  5,
		SIX:   6,
	})

	// Use the dice value
	fmt.Println(dicevalues.V().FOUR) // 4

	// Verify dice value
	fmt.Println(dicevalues.IsValidValue(dicevalues.V().ONE)) // true
	fmt.Println(dicevalues.IsValidValue(dicevalue(7)))       // false

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

	// get all the values as a slice (unordered)
	fmt.Println(dicevalues.Values()) // [1 2 3 4 5 6]

	// get all the keys as a slice (unordered)
	fmt.Println(dicevalues.Keys()) // [ONE TWO THREE FOUR FIVE SIX]

}
