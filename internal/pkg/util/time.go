package util

import (
	"strconv"
	"time"
)

const (
	FMT_DATE_TIME    = "2006-01-02 15:04:05"
	FMT_DATE         = "2006-01-02"
	FMT_TIME         = "15:04:05"
	FMT_DATE_TIME_CN = "2006年01月02日 15时04分05秒"
	FMT_DATE_CN      = "2006年01月02日"
	FMT_TIME_CN      = "15时04分05秒"
)

// NowUnix 秒时间戳
func NowUnix() int64 {
	return time.Now().Unix()
}

// NowTimestamp 毫秒时间戳
func NowTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

// Timestamp 毫秒时间戳
func Timestamp(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// TimeFromUnix 秒时间戳转时间
func TimeFromUnix(unix int64) time.Time {
	return time.Unix(unix, 0)
}

// TimeFromTimestamp 毫秒时间戳转时间
func TimeFromTimestamp(timestamp int64) time.Time {
	return time.Unix(0, timestamp*int64(time.Millisecond))
}

// TimeFormat 时间格式化
func TimeFormat(time time.Time, layout string) string {
	return time.Format(layout)
}

// TimeParse 字符串时间转时间类型
func TimeParse(timeStr, layout string) (time.Time, error) {
	return time.Parse(layout, timeStr)
}

// GetDay return yyyyMMdd
func GetDay(time time.Time) int {
	ret, _ := strconv.Atoi(time.Format("20060102"))
	return ret
}

// WithTimeAsStartOfDay 返回指定时间当天的开始时间
func WithTimeAsStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
