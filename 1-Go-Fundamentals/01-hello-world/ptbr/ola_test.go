package main

import "testing"

func TestOla(t *testing.T) {
	got := Ola()
	want := "Olá, mundo!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
