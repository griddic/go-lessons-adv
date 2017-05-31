package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Byte(t *testing.T) {
	assert.Equal(t, "5B", humanize(5))
	assert.Equal(t, "100B", humanize(100))
	assert.Equal(t, "999B", humanize(999))
}

func Test_KByte(t *testing.T) {
	assert.Equal(t, "1.1KB", humanize(1101))
	assert.Equal(t, "1KB", humanize(1000))
}

func Test_MByte(t *testing.T) {
	assert.Equal(t, "2.1MB", humanize(2100000))
}

func Test_Hole(t *testing.T) {
	assert.Equal(t, "1.01KB", humanize(1010))
	assert.Equal(t, "10.1KB", humanize(10199))
	assert.Equal(t, "101KB", humanize(101000))
}

func Test_One(t *testing.T) {
	assert.Equal(t, "1B", humanize(1))
	assert.Equal(t, "10B", humanize(10))
	assert.Equal(t, "100B", humanize(100))
	assert.Equal(t, "1KB", humanize(1000))
	assert.Equal(t, "10KB", humanize(10000))
	assert.Equal(t, "100KB", humanize(100000))
	assert.Equal(t, "1MB", humanize(1000000))
	assert.Equal(t, "10MB", humanize(10000000))
	assert.Equal(t, "100MB", humanize(100000000))
	assert.Equal(t, "1GB", humanize(1000000000))
	assert.Equal(t, "10GB", humanize(10000000000))
	assert.Equal(t, "100GB", humanize(100000000000))
	assert.Equal(t, "1000GB", humanize(1000000000000))
	assert.Equal(t, "10000GB", humanize(10000000000000))
	assert.Equal(t, "100000GB", humanize(100000000000000))
}

func Test_OneTwo(t *testing.T) {
	assert.Equal(t, "1B", humanize(1))
	assert.Equal(t, "12B", humanize(12))
	assert.Equal(t, "120B", humanize(120))
	assert.Equal(t, "1.2KB", humanize(1200))
	assert.Equal(t, "12KB", humanize(12000))
	assert.Equal(t, "120KB", humanize(120000))
	assert.Equal(t, "1.2MB", humanize(1200000))
	assert.Equal(t, "12MB", humanize(12000000))
	assert.Equal(t, "120MB", humanize(120000000))
	assert.Equal(t, "1.2GB", humanize(1200000000))
	assert.Equal(t, "12GB", humanize(12000000000))
	assert.Equal(t, "120GB", humanize(120000000000))
	assert.Equal(t, "1200GB", humanize(1200000000000))
	assert.Equal(t, "12000GB", humanize(12000000000000))
	assert.Equal(t, "120000GB", humanize(120000000000000))
}

func Test_OneTwoThree(t *testing.T) {
	assert.Equal(t, "1B", humanize(1))
	assert.Equal(t, "12B", humanize(12))
	assert.Equal(t, "123B", humanize(123))
	assert.Equal(t, "1.23KB", humanize(1234))
	assert.Equal(t, "12.3KB", humanize(12345))
	assert.Equal(t, "123KB", humanize(123456))
	assert.Equal(t, "1.23MB", humanize(1234567))
	assert.Equal(t, "12.3MB", humanize(12345678))
	assert.Equal(t, "123MB", humanize(123456789))
	assert.Equal(t, "1.23GB", humanize(1234567890))
	assert.Equal(t, "12.3GB", humanize(12345678901))
	assert.Equal(t, "123GB", humanize(123456789012))
	assert.Equal(t, "1234GB", humanize(1234567890123))
	assert.Equal(t, "12345GB", humanize(12345678901234))
	assert.Equal(t, "123456GB", humanize(123456789012345))
}
