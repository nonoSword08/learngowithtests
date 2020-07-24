package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker  // 此处使用一个文件作为database，测试用例中使用临时文件
}

// 解析数据库中的存储的json数据并以Player列表的形式返回所有存储的数据
func (f *FileSystemPlayerStore) GetLeague() League {
	f.database.Seek(0, 0) // 每次读取时seek到文件开头处
	league, _ := NewLeague(f.database)
	return league
}

// 返回某位玩家的数据
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.GetLeague().Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

// 记录某位玩家得分
func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{name, 1})
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}

// 得分列表包装类
type League []Player

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name==name {
			return &l[i]
		}
	}
	return nil
}