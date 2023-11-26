package util

import (
	"strings"
	"testing"
)

func TestRandomPrefixWithString(t *testing.T) {
	prefix := "prefix"
	result := RandomStringWithPrefix(prefix)
	if !strings.HasPrefix(result, prefix) {
		t.Errorf("RandomStringWithPrefix(%s) should have prefix %s", result, prefix)
	}

}
