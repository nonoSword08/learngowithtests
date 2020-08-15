package poker

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// 提醒管理员输入游戏人数
const PlayerPrompt = "Please enter the munber of players: "

// 游戏人数输入错误时的提醒
const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"

// 赢家输入错误时的提醒
const BadWinnerInputMsg = "invalid winner input, expect format of 'PlayerName wins'"

// 游戏客户端，负责游戏的开始、提醒、结束和记录
type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

// CLI构造函数
func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

// 开始游戏
func (cli *CLI) PlayPoker() {
	// 提示用户输入玩家人数
	fmt.Fprint(cli.out, PlayerPrompt)
	// 读取用户输入的人数，转int
	numberOfPlayers, err := strconv.Atoi(cli.readLine())
	if err != nil {
		fmt.Fprint(cli.out, BadPlayerInputErrMsg)
		return
	}
	// 根据玩家人数来提醒下注
	cli.game.Start(numberOfPlayers)
	// 读取用户输入的赢家
	userInput := cli.readLine()
	winner, err := extractWinner(userInput)
	if err != nil {
		fmt.Fprint(cli.out, BadWinnerInputMsg)
		return
	}
	// 记录赢家
	cli.game.Finish(winner)
}

func extractWinner(userInput string) (string, error) {
	if !strings.Contains(userInput, " wins") {
		return "", errors.New(BadWinnerInputMsg)
	}
	return strings.Replace(userInput, " wins", "", 1), nil
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
