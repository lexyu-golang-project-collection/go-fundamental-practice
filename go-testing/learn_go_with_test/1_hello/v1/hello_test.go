package hello

import "testing"

func TestHello(t *testing.T) {

	t.Run("say hello", func(t *testing.T) {
		actual := Hello("Beta Player")

		result := "Hello, Beta Player"

		assertCorrectMessage(t, actual, result)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		actual := Hello("")

		result := "Hello, World"

		assertCorrectMessage(t, actual, result)
	})
}

func assertCorrectMessage(t testing.TB, actual, result string) {
	t.Helper()
	if actual != result {
		t.Errorf("input %q result %q", actual, result)
	}
}
