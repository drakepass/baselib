package basefile

import (
	"time"
)

const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * Nanosecond
	Millisecond               = 1000 * Microsecond
	Second                    = 1000 * Millisecond
	Minute                    = 60 * Second
	Hour                      = 60 * Minute
	Day                       = 24 * Hour
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

//传入北京时间(月日时分秒要采用两位表示)，获取北京时间戳
func Strtotime(beijingtime string) (int64, error) {
	layout := "2006-01-02 15:04:05"
	time.Local, _ = time.LoadLocation("Asia/Chongqing")
	t, err := time.ParseInLocation(layout, beijingtime, time.Local)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

//将时间戳转换格式（2006-01-02 15.04.05）的时间字符串,可根据传入的duration精确到具体单位
func Timeformat(timestamp int64, duration time.Duration) string {
	layout := "2006-01-02 15:04:05"
	switch duration {
	case Minute:
		layout = "2006-01-02 13:04"
	case Hour:
		layout = "2006-01-02 13"
	case Day:
		layout = "2006-01-02"
	}
	return time.Unix(timestamp, 0).Format(layout)
}
