package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to a person", func(t *testing.T) {
		got := getHello("Anuj","")
		want := "Hello, Anuj"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to an empty string", func(t *testing.T) {
		got := getHello("","")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := getHello("Anuj", "Spanish")
		want := "Hola, Anuj"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T){
		got := getHello("Anuj", "French")
		want := "Bonjour, Anuj"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
