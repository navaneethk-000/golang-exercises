package main

import "testing"

func TestAbbreviate(t *testing.T) {
	shortform := Abbreviate("Work From Home")

	if shortform != "WFH" {
		t.Errorf("Expected WFH but got %s\n", shortform)
	}
}
