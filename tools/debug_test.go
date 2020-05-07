package tools_test

import (
	"github.com/strangebuzz/cc/tools"
	"strings"
	"testing"
)

// This is just a test example
func TestPrettyFormat(t *testing.T) {
	testTable := map[string]bool{
		"Yes": true,
	}
	found := tools.PrettyFormat(testTable)
	found = strings.Replace(found, "\n", "", -1)
	found = strings.Replace(found, "\t", "", -1)
	expected := "{\"Yes\": true}"
	if found != expected {
		t.Errorf("PrettyFormat incorrect, expected: '%s' found: '%s' ", expected, found)
	}
}
