package util

import "testing"

func TestGetCurrentTimestamp(t *testing.T) {
	ts := GetCurrentTimestamp()
	if ts == 0 {
		t.Error()
	}
}
