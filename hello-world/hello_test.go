package main

import "testing"

func TestHello(t *testing.T){
	got := Hello("Batata")
	want := "Hello Batata"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}