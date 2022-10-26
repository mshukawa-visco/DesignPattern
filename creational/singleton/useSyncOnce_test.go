package singleton

import (
	"testing"
)

func TestGetDBInfo(t *testing.T) {
	testcases := []struct {
		name     string
		expected string
	}{
		{"PostgreSQL", "PostgreSQL"},
		{"MySQL", "PostgreSQL"},
		{"SQLite", "PostgreSQL"},
	}

	for _, tt := range testcases {
		db := GetDataBase(tt.name)
		if db == nil {
			t.Fatal("GetDataBase(): expected db pointer is not nil")
		}
		if name := db.GetDBInfo(); name != tt.expected {
			t.Errorf("GetDBInfo(): expected=%s, but got=%s", tt.expected, name)
		}
	}
}
