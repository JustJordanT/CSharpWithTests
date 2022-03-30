package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jordan")
	want := "Hello, Jordan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
