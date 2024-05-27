package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(5, 5)
	expected := 10

	if sum != expected {
		t.Errorf("expected '%d' but actual '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 6)
	fmt.Println(sum)
	// Output: 7
}
