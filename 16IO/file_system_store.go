package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)  // 每次读取时seek到文件开头处
	league, _ := NewLeague(f.database)
	return league
}