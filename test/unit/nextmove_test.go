package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanity(t *testing.T) {
	expected := 1.0
	actual := 1.0
	assert.Equalf(t, expected, actual, "TestSanity: expected %f != actual %f",expected,actual)

}
