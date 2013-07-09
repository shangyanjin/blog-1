package db

import (
	"testing"
)

func TestFields(t *testing.T) {
	input := []string{"one", "two", "three"}
	expected := "one, two, three"
	output := Fields(input...)
	if output != expected {
		t.Errorf("Expecting: %v\nGot: %v", expected, output)
	}
}
