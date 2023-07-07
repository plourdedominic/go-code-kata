package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnTrueFunction(t *testing.T) {
	assert.True(t, returnTrue())
}
