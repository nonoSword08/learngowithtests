package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increamting the conuter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertConuter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantCount)

		for i := 0; i < wantCount; i++ {
			go func(wg *sync.WaitGroup) {
				counter.Inc()
				wg.Done()
			}(&wg)
		}
		wg.Wait()

		assertConuter(t, counter, wantCount)
	})
}

func assertConuter(t *testing.T, got Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d want %d", got.Value(), want)
	}
}
