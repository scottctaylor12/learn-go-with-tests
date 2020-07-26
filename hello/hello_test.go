package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Scott", "")
		want := "Hello, Scott"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Scott", "French")
		want := "Bonjour, Scott"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in German", func(t *testing.T) {
		got := Hello("Scott", "German")
		want := "Hallo, Scott"
		assertCorrectMessage(t, got, want)
	})
}

func ExampleHello() {
	greet := Hello("Scott", "German")
	fmt.Println(greet)
	//Output: Hallo, Scott
}
