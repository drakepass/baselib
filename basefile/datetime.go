package drakelib

import "time"

const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * Nanosecond
	Millisecond               = 1000 * Microsecond
	Second                    = 1000 * Millisecond
	Minute                    = 60 * Second
	Hour                      = 60 * Minute
)

//计时器，进程内记录多个断点执行时间
type TimeCounter struct {
	Startingtime time.Time
	TimeQueue    []time.Duration
}

func GetMicroTimestamp() int64 {
	return time.Now().UnixNano() / int64(Microsecond)
}
func GetMilliTimestamp() int64 {
	return time.Now().UnixNano() / int64(Millisecond)
}
func NewTimeCounter() *TimeCounter {
	return &TimeCounter{
		time.Time{},
		[]time.Duration{},
	}
}
func (t *TimeCounter) Timing() {
	t.Startingtime = time.Now()
}
func (t *TimeCounter) TmpPoint() {
	timePoint := time.Now().Sub(t.Startingtime)
	t.TimeQueue = append(t.TimeQueue, timePoint)
}
func (t *TimeCounter) GetCountTime() []time.Duration {
	return t.TimeQueue
}
