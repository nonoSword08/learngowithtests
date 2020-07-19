package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// 测试get玩家得分列表
func TestLeague(t *testing.T) {

	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}
	
		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)
	
		request := newLeagueRequest()
		response := httptest.NewRecorder()
	
		server.ServeHTTP(response, request)
	
		got := getLeagueFromResponse(t, response.Body)
		assertStatus(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
	})
}
