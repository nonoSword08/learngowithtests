package racer

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))


	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
	}))
	
	slowURL := slowServer.URL
	fasrURl := fastServer.URL

	fmt.Print(slowURL)

	want := fasrURl
	got := Racer(slowURL, fasrURl)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	slowServer.Close()
	fastServer.Close()
}
