package poker

import (
	"fmt"
	"os"
	"time"
)

// 盲注提醒接口，在duration时间后，提醒盲注数目(打印一段话)
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

// 将函数包装成类来实现接口
type BlindAlerterFunc func(duration time.Duration, amount int)

func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	a(duration, amount)
}

func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}
