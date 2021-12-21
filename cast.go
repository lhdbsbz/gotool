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
