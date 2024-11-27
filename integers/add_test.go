package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Adder(2, 2)
	wants := 4

	AssertCorrectMessage(t, got, wants)
}

func ExampleAdder() {
	sum := Adder(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func AssertCorrectMessage(t *testing.T, got, wants int) {
	t.Helper()
	if got != wants {
		t.Errorf("expected '%d' but got '%d'", wants, got)
	}
}
