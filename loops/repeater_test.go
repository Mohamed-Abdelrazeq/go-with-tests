package loops

import "testing"

func TestReaper(t *testing.T) {
	got := Repeater("a", 5)
	wants := "aaaaa"

	if got != wants {
		t.Errorf("expected '%s' but got '%s'", wants, got)
	}
}

func BenchmarkRepeater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeater("a", 5)
	}
}
