package scs

import (
	"strings"
	"testing"
)

// TestGetChunks tests getChunks function.
func TestGetChunks(t *testing.T) {
	var expected = []string{"russian", "warship", "go", "fuck", "yourself"}

	chunks := getChunks("Russian warship, go fuck yourself!")
	for i, r := range chunks {
		if e := expected[i]; e != r {
			t.Errorf("expected %s but %s", e, r)
		}
	}

	if len(chunks) != len(expected) {
		t.Error("incorrect getChunk's result")
	}
}

// TestToUnited tests toUnited function.
func TestToUnited(t *testing.T) {
	var expected = "russianWarshipGoFuckYourself"

	// The camelCase
	result := toUnited("Russian warship, go fuck yourself!", true)
	if result != expected {
		t.Errorf("expected %s but %s", expected, result)
	}

	// The PascalCase
	result = toUnited("Russian warship, go fuck yourself!", false)
	if result != strings.Title(expected) {
		t.Errorf("expected %s but %s", strings.Title(expected), result)
	}
}

// TestToSeparate tests toSeparate function.
func TestToSeparate(t *testing.T) {
	var expected = "russian:warship:go:fuck:yourself"

	result := toSeparate("Russian warship, go fuck yourself!", ":")
	if result != expected {
		t.Errorf("expected %s but %s", expected, result)
	}
}
