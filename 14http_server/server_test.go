package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayer(t *testing.T) {
	t.Run("returm pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/player/Pepper", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
