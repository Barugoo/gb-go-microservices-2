package main

import "testing"

func TestAverage(t *testing.T) {
	expected := 4
	result := Average(4, 4, 4, 4)
	if expected != result {
		t.Error("failed")
	}
}
