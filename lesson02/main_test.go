package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Bytes(t *testing.T) {
	assert.Equal(t, "5b", FormatSize(5), "Все размеры до 1000 байт показывать в байтах")
	assert.Equal(t, "345b", FormatSize(345), "Все размеры до 1000 байт показывать в байтах")
}

func Test_Kilobytes(t *testing.T) {
	assert.Equal(t, "1.1Kb", FormatSize(1167))
	assert.Equal(t, "1Kb", FormatSize(1000))
	assert.Equal(t, "136Kb", FormatSize(136689))
}
