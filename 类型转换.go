package gotool

import (
	"github.com/shopspring/decimal"
	"reflect"
	"time"
)

// =========================转字符串=========================
func TimeToString[T *time.Time | time.Time](in T) string {
	value := reflect.ValueOf(in)
	var t time.Time
	switch value.Kind() {
	case reflect.Ptr:
		if value.IsNil() {
			return ""
		}
		t = value.Elem().Interface().(time.Time)
	case reflect.Struct:
		t = value.Interface().(time.Time)
	default:
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

// =========================时间转秒=========================
func ToSecondNil[T *time.Time | time.Time](in T) *int64 {
	second := ToSecond(in)
	return &second
}

func ToSecond[T *time.Time | time.Time](in T) int64 {
	value := reflect.ValueOf(in)
	var t time.Time
	switch value.Kind() {
	case reflect.Ptr:
		if value.IsNil() {
			return 0
		}
		t = value.Elem().Interface().(time.Time)
	case reflect.Struct:
		t = value.Interface().(time.Time)
	default:
		return 0
	}
	return t.UnixNano() / int64(time.Second)
}

// =========================时间转毫秒=========================
func ToMillisecondNil[T *time.Time | time.Time](in T) *int64 {
	millisecond := ToMillisecond(in)
	return &millisecond
}
func ToMillisecond[T *time.Time | time.Time](in T) int64 {
	value := reflect.ValueOf(in)
	var t time.Time
	switch value.Kind() {
	case reflect.Ptr:
		if value.IsNil() {
			return 0
		}
		t = value.Elem().Interface().(time.Time)
	case reflect.Struct:
		t = value.Interface().(time.Time)
	default:
		return 0
	}
	return t.UnixNano() / int64(time.Millisecond)
}

// =========================毫秒转时间=========================
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

// =========================秒转时间=========================
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

// =========================浮点数工具截取小数=========================
func DecimalToFloat64(amount decimal.Decimal) float64 {
	result, _ := amount.Float64()
	return result
}

// 向下取精度 precision为精度
func FloorDecimal(amount decimal.Decimal, precision int32) decimal.Decimal {
	toolNum := decimal.New(1, precision)
	return amount.Mul(toolNum).Floor().Div(toolNum)
}

// 向上取精度 precision为精度
func CeilDecimal(amount decimal.Decimal, precision int32) decimal.Decimal {
	toolNum := decimal.New(1, precision)
	return amount.Mul(toolNum).Ceil().Div(toolNum)
}
