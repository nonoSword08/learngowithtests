package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// 测试get玩家得分列表
func TestLeague(t *testing.T) {
	store := StubPlayerStore{}
	server := &PlayerServer{store: &store}

	t.Run("it return 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})
}
