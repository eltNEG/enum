package enum

import (
	"encoding/json"
	"errors"
	"log"
)

type Enum[T any, E comparable] struct {
	v        *T
	valueKey map[E]string
}

type EnumValue interface {
	~string | ~int64 | ~float64 | ~uint64 | ~int | ~uint | ~int8 | ~uint8 | ~int16 | ~uint16 | ~int32 | ~uint32
}

// New makes and return a new enum or returns an error
func New[E EnumValue, T any](construct T) (*Enum[T, E], error) {
	kvBytes, err := json.Marshal(construct)
	if err != nil {
		return nil, err
	}

	kv := map[string]E{}
	err = json.Unmarshal(kvBytes, &kv)
	if err != nil {
		return nil, err
	}

	vk := map[E]string{}
	for k, v := range kv {
		if _, ok := vk[v]; ok {
			return nil, errors.New("duplicated enum value")
		}
		vk[v] = k
	}
	return &Enum[T, E]{&construct, vk}, nil
}

// MustNew makes and returns a new enum or panics
func MustNew[E EnumValue, T any](construct T) *Enum[T, E] {
	e, err := New[E, T](construct)
	if err != nil {
		panic(err)
	}
	return e
}

// IsValidValue checks if the given value is a valid enum value
func (e *Enum[T, E]) IsValidValue(value E) (ok bool) {
	_, ok = e.valueKey[value]
	return ok
}

// IsValidStringKey checks if a string is a valid enum key
func (e *Enum[T, E]) IsValidStringKey(key string) (ok bool) {
	for _, k := range e.valueKey {
		if k == key {
			return true
		}
	}
	return false
}

// GetKeyWithValue gets and returns an enum key using the value
func (e *Enum[T, E]) GetKeyWithValue(value E) (key string, ok bool) {
	key, ok = e.valueKey[value]
	return key, ok
}

// MustGetKeyWithValue gets and returns an enum key using the value
func (e *Enum[T, E]) MustGetKeyWithValue(value E) (key string) {
	key, ok := e.valueKey[value]
	if !ok {
		log.Fatalf("invalid enum value: %v", value)
		return
	}
	return key
}

// GetValueWithStringKey gets and returns an enum value using a string key
func (e *Enum[T, E]) GetValueWithStringKey(key string) (value E, ok bool) {
	for v, k := range e.valueKey {
		if k == key {
			return v, true
		}
	}
	return value, false
}

// MustGetValueWithStringKey gets and returns an enum value using a string key
func (e *Enum[T, E]) MustGetValueWithStringKey(key string) (value E) {
	value, ok := e.GetValueWithStringKey(key)
	if !ok {
		log.Fatalf("invalid enum string: %s", key)
		return
	}
	return value
}

// V return the enum values
func (e *Enum[T, E]) V() T {
	return *e.v
}
