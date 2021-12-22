package prom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleAdd(t *testing.T) {
	Add(1)

	c := Get()

	assert.Equal(t, c, 1)
}
