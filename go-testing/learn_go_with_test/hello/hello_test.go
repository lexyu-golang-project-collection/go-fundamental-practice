package hello

import "testing"

func TestHello(t *testing.T) {
	actual := Hello()

	result := "Hello, world"

	if actual != result {
		t.Errorf("input %q result %q", actual, result)
	}

}
