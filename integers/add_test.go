package integers

import "testing"

func TestAdder(t *testing.T) {
	got := Adder(2, 2)
	wants := 4

	AssertCorrectMessage(t, got, wants)
}

func AssertCorrectMessage(t *testing.T, got, wants int) {
	t.Helper()
	if got != wants {
		t.Errorf("expected '%d' but got '%d'", wants, got)
	}
}
