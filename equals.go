package gotool

import (
	"math"
)

// 比较浮点数金额是否相等
func FloatMoneyEquals(f1, f2 float64, precision ...int) bool {
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
