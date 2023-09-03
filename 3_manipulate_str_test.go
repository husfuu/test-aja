package main

import "testing"

func TestManipulateName(t *testing.T) {
	inputName := "dani Maulana"
	expRes := "d4a2nimul"
	got := GetManipulateName(inputName)

	if got != expRes {
		t.Errorf("Expected to be '%s', got %s", expRes, got)
	}
}
