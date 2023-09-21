package gotool

func VersionA_BiggerThanOrEqual_VersionB(versionA, versionB string) bool {
	if versionA >= versionB {
		return true
	}
	return false
}
