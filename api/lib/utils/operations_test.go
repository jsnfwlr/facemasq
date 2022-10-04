package utils

import "testing"

func TestTernary(t *testing.T) {
	input := "true"
	trueOrFalse := Ternary(input == "true", "true", "false").(string)
	if trueOrFalse != "true" {
		t.Error("Ternary function does not work as expected")
	}
	if trueOrFalse == "false" {
		t.Error("Ternary function does not work as expected")
	}

	input = "false"
	trueOrFalse = Ternary(input == "true", "true", "false").(string)
	if trueOrFalse != "false" {
		t.Error("Ternary function does not work as expected")
	}
	if trueOrFalse == "true" {
		t.Error("Ternary function does not work as expected")
	}
}
