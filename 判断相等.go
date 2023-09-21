package gotool

import (
	"fmt"
	"math"
)

// 在前precision位比较两个浮点数是否相等
func FloatMoneyEquals(f1, f2 float64, precision ...int) bool {
	if len(precision) > 1 {
		panic(fmt.Errorf("小数点位数只能传1个或者不传"))
	}
	var MIN float64
	if len(precision) == 0 {
		return f1 == f2
	} else {
		if precision[0] <= 0 {
			return f1 == f2
		} else {
			MIN = math.Pow(0.1, float64(precision[0]))
			if f1 > f2 {
				return f1-f2 < MIN
			} else {
				return f2-f1 < MIN
			}
		}
	}
}
