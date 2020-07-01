package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("csj")
	want := "hello, csj"
	if got != want{
		t.Errorf("got %q wan %q", got ,want)
	}
}