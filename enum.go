package enum

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
)

type Enum[T any, E comparable] struct {
	v        *T
	valueKey map[E]string
}

type SimpleEnum[T any, E comparable] struct {
	Enum[T, E]
}

type EnumValue interface {
	~string | ~int64 | ~float64 | ~uint64 | ~int | ~uint | ~int8 | ~uint8 | ~int16 | ~uint16 | ~int32 | ~uint32
}

type MakeEnumValue interface {
	~uint8
}

// New makes and return a new custom enum or returns an error
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

// Make makes a simple enum
func Make[E MakeEnumValue, T any](construct T) *SimpleEnum[T, E] {
	rc := reflect.ValueOf(construct).Type()
	n := rc.NumField()

	if n > 255 {
		panic("enum values must be less than 256")
	}

	vk := map[E]string{}

	for i := 0; i < n; i++ {
		j := E(i)
		reflect.ValueOf(&construct).Elem().Field(i).Set(reflect.ValueOf(j))
		vk[j] = rc.Field(i).Name
	}

	return &SimpleEnum[T, E]{Enum[T, E]{&construct, vk}}
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

func (e *SimpleEnum[T, E]) GetKeyAtIndex(value E) (string, bool) {
	k, ok := e.valueKey[E(value)]
	return k, ok
}

// V return the enum values
func (e *Enum[T, E]) V() T {
	return *e.v
}

func (e *Enum[T, E]) Keys() []string {
	keys := make([]string, len(e.valueKey))
	i := 0
	for k := range e.valueKey {
		keys[i] = e.MustGetKeyWithValue(k)
		i++
	}
	return keys
}

func (e *Enum[T, E]) Values() []E {
	values := make([]E, len(e.valueKey))
	i := 0
	for k := range e.valueKey {
		values[i] = k
		i++
	}
	return values
}
