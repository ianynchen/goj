package util

import (
	"testing"

	"github.com/issue9/assert"
)

func TestParseInt8(t *testing.T) {

	content := []byte{0xCA}
	val, err := ParseUint8(content)

	assert.True(t, err == nil)
	assert.Equal(t, val, uint8(0xCA))
}

func TestParseInt16(t *testing.T) {

	content := []byte{0xCA, 0xFE}
	val, err := ParseUint16(content)

	assert.True(t, err == nil)
	assert.Equal(t, val, uint16(0xCAFE))
}

func TestParseInt32(t *testing.T) {

	content := []byte{0xCA, 0xFE, 0xBA, 0xBE}
	val, err := ParseUint32(content)

	assert.True(t, err == nil)
	assert.Equal(t, val, uint32(0xCAFEBABE))
}
