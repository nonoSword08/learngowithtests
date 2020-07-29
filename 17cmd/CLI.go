package poker

import (
	"bufio"
	"io"
	"strings"
)


type CLI struct {
	playerStore PlayerStore
	in io.Reader
}

func (cli *CLI) PlayPoker() {
	reader := bufio.NewScanner(cli.in)
	reader.Scan()
	cli.playerStore.RecordWin(extraWinner(reader.Text()))
}

func extraWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}