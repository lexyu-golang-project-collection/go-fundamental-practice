package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCopySrcToDest(t *testing.T) {

	const input = "TEST"
	src := strings.NewReader(input)
	dest := bytes.NewBuffer(make([]byte, 0))
	err := copySrcToDest(src, dest)
	if err != nil {
		t.FailNow()
	}

	result := dest.String()
	if result != input {
		t.Errorf("expected: %s, got: %s", input, result)
	}
}
