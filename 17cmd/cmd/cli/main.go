package main

import (
	"fmt"
	"os"
	poker "sgwt/17cmd"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	db, _ := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	store, _ := poker.NewFileSystemPlayerStore(db)
	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
