package builder

import (
	"testing"
)

func TestBuildHouse(t *testing.T) {
	testcases := []struct {
		objectType string
		expected   string
	}{
		{"Int", "1"},
		{"String", "a"},
	}

	for _, tt := range testcases {
		director := NewDirector(NewBuilder(tt.objectType))
		if output := director.BuildObject(); output != tt.expected {
			t.Errorf("wrong BuildObject(). expected=%s, but got=%s", tt.expected, output)
		}
	}
}
