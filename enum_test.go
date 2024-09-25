package enum_test

import (
	"slices"
	"testing"

	"github.com/eltNEG/enum"
)

func TestNewEnum(t *testing.T) {
	type testenum1 string
	testenumvalue1 := "testenumvalue1"
	testenumvalue2 := "testenumvalue2"

	testenums, err := enum.New[testenum1](struct {
		Key1 testenum1
		Key2 testenum1
	}{
		Key1: "testenumvalue1",
		Key2: "testenumvalue2",
	})

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
		return
	}

	if testenums.V().Key1 != testenum1(testenumvalue1) {
		t.Errorf("Expected: %s, got %s", testenumvalue1, testenums.V().Key1)
	}

	if testenums.V().Key2 != testenum1(testenumvalue2) {
		t.Errorf("Expected: %s, got %s", testenumvalue2, testenums.V().Key2)
	}

	type testenum2 float64
	testenumvalue3 := 3.4
	testenumvalue4 := 5.3

	testenums2, err := enum.New[testenum2](struct {
		Key1 testenum2
		Key2 testenum2
	}{
		Key1: 3.4,
		Key2: 5.3,
	})

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
		return
	}

	if testenums2.V().Key1 != testenum2(testenumvalue3) {
		t.Errorf("Expected: %v, got %v", testenumvalue3, testenums.V().Key1)
	}

	if testenums2.V().Key2 != testenum2(testenumvalue4) {
		t.Errorf("Expected: %v, got %v", testenumvalue4, testenums2.V().Key2)
	}

	keys := testenums.Keys()
	slices.Sort(keys)

	if !slices.Equal(keys, []string{"Key1", "Key2"}) {
		t.Errorf("Expected keys to be [Key1, Key2] but got %v", testenums2.Keys())
	}

	values := testenums2.Values()
	slices.Sort(values)

	if !slices.Equal(values, []testenum2{3.4, 5.3}) {
		t.Errorf("Expected values to be [3.4, 5.3] but got %v", testenums2.Values())
	}
}

func TestMakeEnum(t *testing.T) {
	type testenum1 uint8
	testenumvalue1 := 0
	testenumvalue2 := 2
	testenumvalue3 := 3

	testenums := enum.Make[testenum1](struct {
		Z,
		W,
		A,
		P,
		B,
		C,
		D,
		E testenum1
	}{})

	if testenums.V().Z != testenum1(testenumvalue1) {
		t.Errorf("Expected: %d, got %d", testenumvalue1, testenums.V().Z)
	}

	if testenums.V().A != testenum1(testenumvalue2) {
		t.Errorf("Expected: %d, got %d", testenumvalue2, testenums.V().A)
	}

	if testenums.V().P != testenum1(testenumvalue3) {
		t.Errorf("Expected: %d, got %d", testenumvalue2, testenums.V().P)
	}

	if testenums.V().A == testenums.V().Z || testenums.V().A == testenums.V().P || testenums.V().Z == testenums.V().P {
		t.Errorf("Expected different values for enum keys")
	}

	if key, _ := testenums.GetKeyAtIndex(0); key != "Z" {
		t.Errorf("Expected: Z, got %s", key)
	}

	if key, _ := testenums.GetKeyAtIndex(2); key != "A" {
		t.Errorf("Expected: A, got %s", key)
	}

	if key, _ := testenums.GetKeyAtIndex(3); key != "P" {
		t.Errorf("Expected: P, got %s", key)
	}

	if !testenums.IsValidValue(testenum1(0)) {
		t.Error("Expected valid value for 0")
	}

	if testenums.IsValidValue(testenum1(100)) {
		t.Error("Expected invalid value for 100")
	}

	keys := testenums.Keys()
	slices.Sort(keys)

	if !slices.Equal(keys, []string{"A", "B", "C", "D", "E", "P", "W", "Z"}) {
		t.Errorf("Expected keys to be [A, B, C, D, E, P, W, Z] but got %v", testenums.Keys())
	}

	values := testenums.Values()
	slices.Sort(values)

	if !slices.Equal(values, []testenum1{0, 1, 2, 3, 4, 5, 6, 7}) {
		t.Errorf("Expected values to be [0, 1, 2, 3, 4, 5, 6, 7] but got %v", testenums.Values())
	}
}

func TestEnumMethods(t *testing.T) {
	type sampleEnum string

	sampleEnums := enum.MustNew[sampleEnum](struct {
		VALUE1 sampleEnum
		VALUE2 sampleEnum
	}{
		VALUE1: "valueone",
		VALUE2: "v2",
	})

	if !sampleEnums.IsValidValue("valueone") {
		t.Error("IsValidValue: expected valid value for valueone")
	}

	if !sampleEnums.IsValidValue("v2") {
		t.Error("IsValidValue: expected valid value for v2")
	}

	if sampleEnums.IsValidValue("V2") {
		t.Error("IsValidValue: expected invalid value for V2")
	}

	if sampleEnums.IsValidStringKey("V2") {
		t.Error("IsValidStringKey: expected key V2 to be invalid")
	}

	if !sampleEnums.IsValidStringKey("VALUE2") {
		t.Error("IsValidStringKey: expected a valid value for VALUE2")
	}

	key, _ := sampleEnums.GetKeyWithValue("v2")
	expectedKey := "VALUE2"

	if key != expectedKey {
		t.Errorf("GetKeyWithValue: expected %s but got %s", expectedKey, key)
	}

	value, _ := sampleEnums.GetValueWithStringKey("VALUE1")
	expectedValue := sampleEnum("valueone")

	if value != expectedValue {
		t.Errorf("GetValueWithStringKey: expected %s but got %s", expectedValue, value)
	}

	_ = sampleEnums.MustGetKeyWithValue("valueone")
	_ = sampleEnums.MustGetValueWithStringKey("VALUE1")

}
