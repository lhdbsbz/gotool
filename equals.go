package gotool

// FloatMoneyEquals 比较浮点数金额是否相等
func FloatMoneyEquals(f1, f2 float64, precision ...float64) bool {
	var MIN = 0.00001
	if len(precision) == 0 {
		MIN = precision[0]
	}
	if f1 > f2 {
		return f1-f2 < MIN
	} else {
		return f2-f1 < MIN
	}
}
