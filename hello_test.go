package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Mohit")
		want := "Hello Mohit"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {

		got := Hello("")
		want := "Hello World"

		if got != want {
			t.Errorf("got %q want %q", got, want)

		}

		assertCorrectMessage(t, got, want)

	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // to tell the compiler to display error on the above lines not here
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
