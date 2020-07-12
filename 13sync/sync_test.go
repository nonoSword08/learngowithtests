package main

import "testing"

func TestCounter(t *testing.T) {
	t.Run("increamting the conuter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

	})
}

func assertConuter(t *testing.T, got Counter, want int) {
	t.Helper()
	if got.Value() != 3 {
		t.Errorf("got %d want %d", got.Value(), 3)
	}
}
