package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

// 解析数据库中的存储的json数据并以Player列表的形式返回所有存储的数据
func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0) // 每次读取时seek到文件开头处
	league, _ := NewLeague(f.database)
	return league
}

// 返回某位玩家的数据
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int
	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}
	return wins
}
