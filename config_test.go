package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadingFile(t *testing.T) {
	Flush()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Loading correct and supported files must not raise panic")
		}
	}()

	LoadFile("appconfig", "yml", "test/")
	LoadFile("appconfig", "json", "test/")
	LoadFile("appconfig", "env", "test/")
}

func TestLoadingUnsupportedFile(t *testing.T) {
	Flush()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Loading unsupported file must raise panic")
		}
	}()

	LoadFile("appconfig", "jsonn", "test/")
}

func TestLoadingFileFromWrongPath(t *testing.T) {
	Flush()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Loading file from wrong path must raise panic")
		}
	}()

	LoadFile("appconfig", "yml", "tests/")
}

func TestLoadingWrongFileName(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Loading wrong file name must raise panic")
		}
	}()

	LoadFile("appConfig", "env", "test/")
}

func TestGet(t *testing.T) {
	v := struct {
		IntValue   int
		StrValue   string
		FloatValue float32
	}{
		IntValue:   1,
		StrValue:   "test",
		FloatValue: 1.12345,
	}

	Set("get", v)
	assert.Equal(t, v, Get("get"))
}

func TestGetString(t *testing.T) {
	Set("getString", "Test GetString")
	assert.Equal(t, "Test GetString", GetString("getString"))
}

func TestGetBool(t *testing.T) {
	Set("testbool", true)
	assert.True(t, GetBool("testbool"))

	Set("testbool", false)
	assert.False(t, GetBool("testbool"))
}
