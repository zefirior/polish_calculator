package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorageStoreValue(t *testing.T) {
	s := NewStorage()
	expected := 100

	s.Set("a", expected)
	actual := s.Get("a")
	assert.Equalf(t, expected, actual, "Expected: %q; Got: %q", expected, actual)
}

func TestFromString(t *testing.T) {
	s := FromString("100 10 1")
	assert.Equal(t, s.Get("a"), 100)
	assert.Equal(t, s.Get("b"), 10)
	assert.Equal(t, s.Get("c"), 1)
}