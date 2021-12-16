package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestThis(t *testing.T) {
	var isTrue = true
	assert.Check(t, isTrue == true, "Should be true all the time.")
}
