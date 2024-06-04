package iteration

import (
	"fmt"
	"testing"
)

func TestReapt(t *testing.T) {

	actual := Repeat("a")
	expected := "aaaaa"

	if expected != actual {
		t.Errorf("expected %q but got %q", expected, actual)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func ExampleRepeat() {
	res := Repeat("c")
	fmt.Println(res)
	// Output: ccccc
}
