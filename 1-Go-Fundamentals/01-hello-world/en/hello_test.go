// ! Lets create a test for the function Hello()
// ? We separated our domain to better testing

package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Exemplo")
	want := "Hello, Exemplo"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
