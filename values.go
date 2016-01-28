package webapi

import (
	"errors"
	"net/url"
	"strconv"
)

// Values maps a string key to a list of values.
// It is typically used for query parameters and form values.
// Unlike in the http.Header map, the keys in a Values map
// are case-sensitive.
type Values struct {
	url.Values
}

var (
	ErrValueNotFinded = errors.New("value not finded")
)

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (this Values) String(key string) (string, error) {
	if this.Values == nil {
		return "", ErrValueNotFinded
	}

	vs, ok := this.Values[key]

	if !ok || len(vs) == 0 {
		return "", ErrValueNotFinded
	}

	return vs[0], nil
}

func (this Values) Int64(key string) (int64, error) {
	if this.Values == nil {
		return 0, ErrValueNotFinded
	}

	vs, ok := this.Values[key]

	if !ok || len(vs) == 0 {
		return 0, ErrValueNotFinded
	}

	return strconv.ParseInt(vs[0], 10, 64)
}

func (this Values) Int32(key string) (int32, error) {
	val64, err := this.Int64(key)

	return int32(val64), err
}

func (this Values) UInt64(key string) (uint64, error) {
	if this.Values == nil {
		return 0, ErrValueNotFinded
	}

	vs, ok := this.Values[key]

	if !ok || len(vs) == 0 {
		return 0, ErrValueNotFinded
	}

	return strconv.ParseUint(vs[0], 10, 64)
}

func (this Values) UInt16(key string) (uint16, error) {
	val64, err := this.UInt64(key)

	return uint16(val64), err
}

/*
// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (v Values) Get(key string) string {
	if v == nil {
		return ""
	}
	vs, ok := v[key]
	if !ok || len(vs) == 0 {
		return ""
	}
	return vs[0]
}

// Set sets the key to value. It replaces any existing
// values.
func (v Values) Set(key, value string) {
	v[key] = []string{value}
}

// Add adds the value to key. It appends to any existing
// values associated with key.
func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}

// Del deletes the values associated with key.
func (v Values) Del(key string) {
	delete(v, key)
}
*/
