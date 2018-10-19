package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseString(t *testing.T) {
	res := ""

	err := ParseString("Foo", &res)

	assert.Equal(t, "Foo", res, "It should return Foo")
	assert.Nil(t, err)
}

func TestParseBool(t *testing.T) {
	res := false

	err := ParseBool("true", &res)

	assert.Equal(t, true, res, "It should return bool true")
	assert.Nil(t, err)
}

func TestParseInt32(t *testing.T) {
	var res int32
	var expected int32 = 123

	err := ParseInt32("123", &res)

	assert.Equal(t, expected, res, "It should return type int32")
	assert.Nil(t, err)
}

func TestParseInt(t *testing.T) {
	var res int
	var expected = 123

	err := ParseInt("123", &res)

	assert.Equal(t, expected, res, "It should return type int")
	assert.Nil(t, err)
}
