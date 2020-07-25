package main

import (
	"log"
	"net/http"
	"os"
)

const dbFilename = "game.db.json"

func main() {
    // 创建一个文件用以保存玩家数据
    db, err := os.OpenFile(dbFilename, os.O_RDWR|os.O_CREATE, 0666)
    if err != nil {
        log.Fatalf("open file fail %s", err)
    }
    server := NewPlayerServer(NewFileSystemPlayerStore(db))

    if err := http.ListenAndServe(":5000", server); err != nil {
        log.Fatalf("could not listen on port 5000 %v", err)
    }
}
