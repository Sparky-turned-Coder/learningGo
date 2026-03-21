package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris!"

	if got != want {
		t.Errorf("got %q and we wanted %q", got, want)
	}
}
