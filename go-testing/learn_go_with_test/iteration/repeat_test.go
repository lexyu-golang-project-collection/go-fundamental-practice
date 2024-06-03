package iteration

import "testing"

func TestReapt(t *testing.T) {

	actual := Repeat("a")
	expected := "aaaaa"

	if expected != actual {
		t.Errorf("expected %q but got %q", expected, actual)
	}

}
