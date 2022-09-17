package main

import (
	"testing"
)

func TestReverse_app_end(t *testing.T) {
	got := "cezar"
	want := reverseAppEnd(got)

	if want != "razec" {
		t.Errorf("got %q , expected %q", got, want)
	}
}

func TestReverseByRune(t *testing.T) {
	got := "cezar"
	want := reverseByRune(got)

	if want != "razec" {
		t.Errorf("got %q , expected %q", got, want)
	}
}
