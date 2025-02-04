package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Mohit", "")
		want := "Hello Mohit"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {

		got := Hello("", "")
		want := "Hello World"

		if got != want {
			t.Errorf("got %q want %q", got, want)

		}

		assertCorrectMessage(t, got, want)

	})

	t.Run("Saying Hola for Spanish", func(t *testing.T) {

		got := Hello("Mohit", "Spanish")
		want := "Hola, Mohit"
		assertCorrectMessage(t, got, want)

	})

	t.Run("Saying Bonnjour for French", func(t *testing.T) {

		got := Hello("Mohit", "French")
		want := "Bonjour, Mohit"
		assertCorrectMessage(t, got, want)

	})
	t.Run("Saying Namaskar for Hindi", func(t *testing.T) {

		got := Hello("Mohit", "Hindi")
		want := "Namaskar, Mohit"
		assertCorrectMessage(t, got, want)

	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // to tell the compiler to display error on the above lines not here
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
