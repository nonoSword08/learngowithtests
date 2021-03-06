package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type FileSystemPlayerStore struct {
	database *json.Encoder // 此处使用一个文件作为database，测试用例中使用临时文件
	league   League        // 临时将数据存储在内存中
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	// 只读取一次文件，将json字符串保存在内存中，只在更新时将其写入文件
	file.Seek(0, 0)

	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil

}

// 解析数据库中的存储的json数据并以Player列表的形式返回所有存储的数据
func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

// 返回某位玩家的数据
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

// 记录某位玩家得分
func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.database.Encode(f.league)
}

// 得分列表包装类
type League []Player

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}
