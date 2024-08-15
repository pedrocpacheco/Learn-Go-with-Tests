// ! Lets create a test for the function Hello()
// ? We separated our domain to better testing

package main

import "testing"

func TestHello(t *testing.T) {
	// assinatura t.Run -> func (t *T) Run(name string, f func(t *T)) bool
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Exemplo")
		want := "Hello, Exemplo"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}
