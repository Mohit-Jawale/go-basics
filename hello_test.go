package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Mohit")
	want := "Hello Mohit"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
