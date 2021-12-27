package gotool

import (
	"github.com/spf13/cast"
	"time"
)

// ToString 转字符串
func ToString(item interface{}) string {
	switch s := item.(type) {
	case nil:
		return ""
	case time.Time:
		// 这是golang诞生的时间 此处传其他参数会报错
		return s.Format("2006-01-02 15:04:05")
	default:
		return cast.ToString(item)
	}
}

// ToMillisecond 时间转毫秒
func ToMillisecond(in time.Time) int64 {
	return *ToMillisecondNil(&in)
}
func ToMillisecondNil(in *time.Time) *int64 {
	if in == nil {
		return nil
	}
	result := in.UnixNano() / int64(time.Millisecond)
	return &result
}

// MillToTime 毫秒转时间
func MillToTime(in int64) time.Time {
	return *MillToTimeNil(&in)
}

func MillToTimeNil(in *int64) *time.Time {
	if in == nil {
		return nil
	}
	unix := time.Unix(0, *in*int64(time.Millisecond))
	return &unix
}

// SecondToTime 秒转时间
func SecondToTime(in int64) time.Time {
	return *SecondToTimeNil(&in)
}

func SecondToTimeNil(in *int64) *time.Time {
	if in == nil {
		return nil
	}
	unix := time.Unix(0, *in*int64(time.Second))
	return &unix
}

// ToSecond 时间转秒
func ToSecond(in time.Time) int64 {
	return *ToSecondNil(&in)
}
func ToSecondNil(in *time.Time) *int64 {
	if in == nil {
		return nil
	}
	result := in.UnixNano() / int64(time.Second)
	return &result
}
