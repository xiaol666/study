package dataStructuresAndAlgorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	hs := NewHashSet()
	hs.Add(3)
	hs.Add(192)
	hs.Add(67)

	assert.Equal(t, true, hs.Contains(3))
	assert.Equal(t, true, hs.Contains(192))
	assert.Equal(t, true, hs.Contains(67))
	assert.Equal(t, false, hs.Contains(9))
}

func TestRemove(t *testing.T) {
	hs := NewHashSet()
	hs.Add(3)
	assert.Equal(t, true, hs.Contains(3))
	hs.Remove(3)
	assert.Equal(t, false, hs.Contains(3))
	hs.Add(9)
	assert.Equal(t, true, hs.Contains(9))
	hs.Remove(3)
	assert.Equal(t, true, hs.Contains(9))
	hs.Remove(9)
	assert.Equal(t, false, hs.Contains(9))
}
