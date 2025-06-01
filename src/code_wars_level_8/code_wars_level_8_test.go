package main

import (
	"testing"
)

func TestEvenOrOddTwo(t *testing.T) {
	result := EvenOrOdd(2)
	want := "Even"
	if result != want {
		t.Errorf("incorrect return")
	}
}

func TestEvenOrOddThree(t *testing.T) {
	result := EvenOrOdd(3)
	want := "Odd"
	if result != want {
		t.Errorf("incorrect return")
	}
}

func TestEvenOrOddNegOne(t *testing.T) {
	result := EvenOrOdd(-1)
	want := "Odd"
	if result != want {
		t.Errorf("incorrect return")
	}
}

func TestEvenOrOddZero(t *testing.T) {
	result := EvenOrOdd(0)
	want := "Even"
	if result != want {
		t.Errorf("incorrect return")
	}
}
