package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Bytes(t *testing.T) {
	assert.Equal(t, "5b", FormatSize(5), "Все размеры до 1000 байт показывать в байтах")
	assert.Equal(t, "345b", FormatSize(345), "Все размеры до 1000 байт показывать в байтах")
	assert.Equal(t, "999b", FormatSize(999), "Все размеры до 1000 байт показывать в байтах")
}

func Test_Kilobytes(t *testing.T) {
	assert.Equal(t, "1.00Kb", FormatSize(1000))
	assert.Equal(t, "1.17Kb", FormatSize(1167))
	assert.Equal(t, "1.30Kb", FormatSize(1300))
	assert.Equal(t, "58.2Kb", FormatSize(58226))
	assert.Equal(t, "137Kb", FormatSize(136689))
	assert.Equal(t, "999Kb", FormatSize(999400))
	assert.Equal(t, "1.00Mb", FormatSize(999500))
	assert.Equal(t, "1.00Mb", FormatSize(999999))
}

func Test_Megabytes(t *testing.T) {
	assert.Equal(t, "1.00Mb", FormatSize(1000000))
	assert.Equal(t, "1.59Mb", FormatSize(1585887))
	assert.Equal(t, "84.4Mb", FormatSize(84369689))
	assert.Equal(t, "562Mb", FormatSize(562029201))
	assert.Equal(t, "999Mb", FormatSize(999324922))
	assert.Equal(t, "1.00Gb", FormatSize(999999999))
}

func Test_Gigabytes(t *testing.T) {
	assert.Equal(t, "1.00Gb", FormatSize(1000000000))
	assert.Equal(t, "6.72Gb", FormatSize(6719459231))
	assert.Equal(t, "94.7Gb", FormatSize(94698830910))
	assert.Equal(t, "100Gb", FormatSize(100000000000))
	assert.Equal(t, "745Gb", FormatSize(745123456789))
	assert.Equal(t, "1.00Tb", FormatSize(999999999999))
}

func Test_Terabytes(t *testing.T) {
	assert.Equal(t, "1.00Tb", FormatSize(1000000000000))
	assert.Equal(t, "3.25Tb", FormatSize(3253920031256))
	assert.Equal(t, "54.4Tb", FormatSize(54368983209010))
	assert.Equal(t, "500Tb", FormatSize(500000000000000))
	assert.Equal(t, "931Tb", FormatSize(931123545267789))
	assert.Equal(t, "1.00Pb", FormatSize(999999999999999))
}

func Test_Petabytes(t *testing.T) {
	assert.Equal(t, "1.00Pb", FormatSize(1000000000000000))
	assert.Equal(t, "4.24Pb", FormatSize(4236459293827734))
	assert.Equal(t, "73.8Pb", FormatSize(73798445332830910))
	assert.Equal(t, "100Pb", FormatSize(100000000000000000))
	assert.Equal(t, "344Pb", FormatSize(344123845156778789))
	assert.Equal(t, "1.00Eb", FormatSize(999999999999999999))
}
