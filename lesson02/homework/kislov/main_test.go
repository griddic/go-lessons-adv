package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Bytes(t *testing.T) {
	assert.Equal(t, "5b", FormatSize(5), "Все размеры до 1024 байт показывать в байтах")
	assert.Equal(t, "345b", FormatSize(345), "Все размеры до 1024 байт показывать в байтах")

}

func Test_Kilobytes(t *testing.T) {
	assert.Equal(t, "1.1KB", FormatSize(1167))
	assert.Equal(t, "1KB", FormatSize(1024))
	assert.Equal(t, "133KB", FormatSize(136689))

}
func Test_Megabytes(t *testing.T) {
	assert.Equal(t, "1MB", FormatSize(1048576))
	assert.Equal(t, "1.5MB", FormatSize(1572864))
	assert.Equal(t, "100MB", FormatSize(104857600))
}
func Test_Gigabytes(t *testing.T) {
	assert.Equal(t, "1GB", FormatSize(1073741824))
	assert.Equal(t, "1.6GB", FormatSize(1610612736))
	assert.Equal(t, "100GB", FormatSize(107374182400))
}