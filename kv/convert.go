package kv

import (
	"encoding/binary"
	"errors"
	"math"
)

var (
	// ErrUnmarshal is returned when unmarshaling to a native type failed.
	ErrUnmarshal = errors.New("unmarshaling failed")
	// ErrMarshal is returned when marshaling from a native type failed.
	ErrMarshal = errors.New("marshaling failed")
)

func marshalBytesInt64(out int64) []byte {
	byt := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(byt, out)
	return byt
}

func marshalBytesUint64(out uint64) []byte {
	byt := make([]byte, binary.MaxVarintLen64)
	binary.PutUvarint(byt, out)
	return byt
}

// CreateUint64 creates a new key as uint64.
func (c *Client) CreateUint64(key string, out uint64) ([]byte, error) {
	return c.driver.Create(c.member, key, marshalBytesUint64(out))
}

// GetUint64 retrieves the marshaled data for key and then converts it to uint64.
func (c *Client) GetUint64(key string) (uint64, []byte, error) {
	out, nonce, err := c.driver.Get(c.member, key)
	if err != nil {
		return 0, nonce, err
	}

	u, left := binary.Uvarint(out)
	if left <= 0 {
		return 0, nonce, ErrUnmarshal
	}

	return u, nonce, nil
}

// SetUint64 sets a uint64 to key.
func (c *Client) SetUint64(key string, out uint64) ([]byte, error) {
	return c.driver.Set(c.member, key, marshalBytesUint64(out))
}

// CASUint64 compares and swaps uint64 values.
func (c *Client) CASUint64(key string, nonce []byte, origValue, value uint64) ([]byte, error) {
	return c.driver.CAS(c.member, key, nonce, marshalBytesUint64(origValue), marshalBytesUint64(value))
}

// CreateInt64 creates a new key as int64.
func (c *Client) CreateInt64(key string, out int64) ([]byte, error) {
	return c.driver.Create(c.member, key, marshalBytesInt64(out))
}

// GetInt64 retrieves the marshaled data for key and then converts it to int64.
func (c *Client) GetInt64(key string) (int64, []byte, error) {
	out, nonce, err := c.driver.Get(c.member, key)
	if err != nil {
		return 0, nonce, err
	}

	v, left := binary.Varint(out)
	if left <= 0 {
		return 0, nonce, ErrUnmarshal
	}

	return v, nonce, nil
}

// SetInt64 sets a int64 to key.
func (c *Client) SetInt64(key string, out int64) ([]byte, error) {
	return c.driver.Set(c.member, key, marshalBytesInt64(out))
}

// CASInt64 compares and swaps uint64 values.
func (c *Client) CASInt64(key string, nonce []byte, origValue, value int64) ([]byte, error) {
	return c.driver.CAS(c.member, key, nonce, marshalBytesInt64(origValue), marshalBytesInt64(value))
}

// CreateString creates a string where there wasn't one before.
func (c *Client) CreateString(key, value string) ([]byte, error) {
	return c.driver.Create(c.member, key, []byte(value))
}

// GetString retrieves the marshaled data for key and then converts it to string.
func (c *Client) GetString(key string) (string, []byte, error) {
	out, nonce, err := c.driver.Get(c.member, key)
	if err != nil {
		return "", nonce, err
	}

	return string(out), nonce, nil
}

// SetString sets a string to a key.
func (c *Client) SetString(key, value string) ([]byte, error) {
	return c.driver.Set(c.member, key, []byte(value))
}

// CASString compares and swaps strings.
func (c *Client) CASString(key string, nonce []byte, origValue, value string) ([]byte, error) {
	return c.driver.CAS(c.member, key, nonce, []byte(origValue), []byte(value))
}

// CreateFloat64 creates a float64 where there wasn't one before.
func (c *Client) CreateFloat64(key string, out float64) ([]byte, error) {
	return c.CreateUint64(key, math.Float64bits(out))
}

// GetFloat64 retrieves the marshaled data for key and then converts it to float64.
func (c *Client) GetFloat64(key string) (float64, []byte, error) {
	u, nonce, err := c.GetUint64(key)
	if err != nil {
		return 0, nonce, err
	}

	return math.Float64frombits(u), nonce, nil
}

// SetFloat64 sets a float64 to a key.
func (c *Client) SetFloat64(key string, value float64) ([]byte, error) {
	return c.SetUint64(key, math.Float64bits(value))
}

// CASFloat64 compares and swaps float64s.
func (c *Client) CASFloat64(key string, nonce []byte, origValue, value float64) ([]byte, error) {
	return c.CASUint64(key, nonce, math.Float64bits(origValue), math.Float64bits(value))
}
