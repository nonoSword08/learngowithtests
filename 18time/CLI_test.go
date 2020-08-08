package poker

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)

// 盲注提醒类
type scheduledAlert struct {
	at     time.Duration
	amount int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

// 测试用盲注提醒用例
type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{duration, amount})
}

// 其他不相关测试用傀儡对象
var dummySpyAlerter = &SpyBlindAlerter{}

// 用于测试CLI输入输出是否正确的傀儡对象
var dummyBlindAlerter = &SpyBlindAlerter{}
var dummyPlayerStore = &StubPlayerStore{}
var dummyStdin = &bytes.Buffer{}
var dummyStdout = &bytes.Buffer{}

func TestCLI(t *testing.T) {

	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &StubPlayerStore{}

		cli := NewCLI(playerStore, in, dummyStdout, dummySpyAlerter)
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &StubPlayerStore{}

		cli := NewCLI(playerStore, in, dummyStdout, dummySpyAlerter)
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Cleo")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := NewCLI(playerStore, in, dummyStdout, blindAlerter)
		cli.PlayPoker()

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduleAlert(t, got, want)

			})
		}
	})

	t.Run("it prompts the user to enter th number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		cli := NewCLI(dummyPlayerStore, dummyStdin, stdout, dummyBlindAlerter)
		cli.PlayPoker()

		got := stdout.String()
		want := PlayerPrompts

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func assertScheduleAlert(t *testing.T, got scheduledAlert, want scheduledAlert) {
	if got != want {
		t.Errorf("got scheduledAlert %d, want %d", got, want)
	}
}
