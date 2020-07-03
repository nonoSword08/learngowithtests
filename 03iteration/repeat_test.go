package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeatd := Repeat("a")
	exepected := "aaaaa"

	if repeatd != exepected {
		t.Errorf("expected %q but got %q", exepected, repeatd)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
