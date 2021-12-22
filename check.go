package gotool

import (
	"fmt"
	"reflect"
	strings "strings"
)

func IsNotEmpty(x interface{}) bool {
	return !IsEmpty(x)
}

func IsEmpty(in interface{}) bool {
	if in == nil {
		return true
	}
	value := reflect.ValueOf(in)
	switch value.Kind() {
	case reflect.String:
		return len(strings.TrimSpace(value.String())) == 0
	case reflect.Bool, reflect.Func, reflect.Invalid, reflect.Uintptr, reflect.UnsafePointer, reflect.Chan:
		panic(fmt.Sprintf("不支持对%v类型调用IsEmpty", value.Kind()))
	case reflect.Slice, reflect.Map:
		return value.Len() == 0
	default:
		return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
	}
}
