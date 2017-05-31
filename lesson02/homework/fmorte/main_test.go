package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Bytes(t *testing.T) {
	assert.Equal(t, "5b", FormatSize(5), "Все размеры до 1000 байт показывать в байтах")
	assert.Equal(t, "345b", FormatSize(345), "Все размеры до 1000 байт показывать в байтах")
}

func Test_Kilobytes(t *testing.T) {
	assert.Equal(t, "1Kb", FormatSize(999), "Все размеры до 1000 килобайт показывать в килобайтах")
	assert.Equal(t, "1.2Kb", FormatSize(1167), "Все размеры до 1000 килобайт показывать в килобайтах")
	assert.Equal(t, "1.1Kb", FormatSize(1067), "Все размеры до 1000 килобайт показывать в килобайтах")
	assert.Equal(t, "136.7Kb", FormatSize(136689), "Все размеры до 1000 килобайт показывать в килобайтах")
}

func Test_Megabytes(t *testing.T) {
	assert.Equal(t, "1Mb", FormatSize(999950), "Все размеры до 1000 килобайт показывать в килобайтах")
	assert.Equal(t, "11.7Mb", FormatSize(11678988), "Все размеры до 1000 килобайт показывать в килобайтах")
	assert.Equal(t, "1Mb", FormatSize(1000067), "Все размеры до 1000 мегабайт показывать в мегабайтах")
	assert.Equal(t, "10Mb", FormatSize(9999999), "Все размеры до 1000 мегабайт показывать в мегабайтах")
}

func Test_Gigabytes(t *testing.T) {
	assert.Equal(t, "1Gb", FormatSize(999950000), "Все размеры до 1000 гигабайт показывать в гигабайтах")
	assert.Equal(t, "1.1Gb", FormatSize(1100006700), "Все размеры до 1000 гигабайт показывать в гигабайтах")
	assert.Equal(t, "91.1Gb", FormatSize(91100006700), "Все размеры до 1000 гигабайт показывать в гигабайтах")
}
