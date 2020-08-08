package poker

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const PlayerPrompts = "Please enter the munber of players: "

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

// 接口实现类
type BlindAlerterFunc func(duration time.Duration, amount int)

func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	a(duration, amount)
}

// 用于构造接口实现类
func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game *Game
}

func NewCLI(store PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI {
	return &CLI{
		in:  bufio.NewScanner(in),
		out: out,
		game: &Game{
			alerter: alerter,
			store:   store,
		},
	}
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func (cli *CLI) PlayPoker() {
	// 提示用户输入玩家人数
	fmt.Fprint(cli.out, PlayerPrompts)
	// 读取用户输入的人数，转int
	numberOfPlayers, _ := strconv.Atoi(cli.readLine())
	// 根据玩家人数来提醒下注
	cli.game.Start(numberOfPlayers)
	// 读取用户输入的赢家
	userInput := cli.readLine()
	winner := extractWinner(userInput)
	// 记录赢家
	cli.game.store.RecordWin(winner)
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

// func (cli *CLI) scheduleBlindAlerts(numberOfPlayers int) {
// 	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

// 	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
// 	blindTime := 0 * time.Second

// 	for _, blind := range blinds {
// 		cli.alerter.ScheduleAlertAt(blindTime, blind)
// 		blindTime = blindTime + blindIncrement
// 	}

// }
