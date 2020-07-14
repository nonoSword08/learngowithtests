package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	server := &PlayerServer{store: &InMemoryPlayerStore{}}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

// 应用用的存储得分结构体
type InMemoryPlayerStore struct {}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}