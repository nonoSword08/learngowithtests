package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T){
		got:= Hello("chris")
		want := "hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'hello world' when an enmpty string is supplied", func(t *testing.T){
		got := Hello("")
		want := "hello world"

		if got != want {
			t.Errorf("got %q want %q", got ,want)
		}
	})
}