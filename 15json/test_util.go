package main

import (
	"testing"
)

// 测试时使用的存储得分结构体
type StubPlayerStore struct {
	scores   map[string]int // 每位玩家的得分
	winCalls []string       // 得分历史
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
