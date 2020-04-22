package main

import "testing"

func TestCalc(t *testing.T) {
	if Calucate(5) != 4 {
		t.Error("Expected 9")
	}
}
func TestCalcPositive(t *testing.T) {
	if Calucate(5) != 7 {
		t.Error("Expected 7")
	}
}
