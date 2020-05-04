package tools

import (
	"strings"
	"testing"
)

// This is just a test example
func TestPrettyFormat(t *testing.T) {
	testTable := map[string]bool{
		"Yes": true,
	}
	found := PrettyFormat(testTable)
	found = strings.Replace(found, "\n", "", -1)
	found = strings.Replace(found, "\t", "", -1)
	expected := "{\"Yes\": true}"
	if found != expected {
		t.Errorf("PrettyFormat incorrect, expected: '%s' found: '%s' ", expected, found)
	}
	PrettyPrint("")
}
