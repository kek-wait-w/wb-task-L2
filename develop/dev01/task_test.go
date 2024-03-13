package main

import "testing"

func TestGetCurrentTime(t *testing.T) {
	_, err := getTimeNow()
	if err != nil {
		t.Fatalf("TEST ERROR: %v", err)
	}
}
