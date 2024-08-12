package enum_test

import (
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
