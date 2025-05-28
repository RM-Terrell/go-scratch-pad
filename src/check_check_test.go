package main

import (
	"testing"
)

func TestCheckCheck(t *testing.T) {
	result := checkCheck()
	want := 5
	if result != want {
		t.Errorf("incorrect return")
	}
}
