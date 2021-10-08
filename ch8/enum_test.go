package ch8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnum(t *testing.T) {
	assert.Equal(t, 1, SUNDAY)
	assert.Equal(t, 2, MONDAY)
	assert.Equal(t, 3, TUESDAY)
	assert.Equal(t, 4, WEDNESDAY)
	assert.Equal(t, 5, THURSDAY)
	assert.Equal(t, 6, FRIDAY)
	assert.Equal(t, 7, SATURDAY)
}
