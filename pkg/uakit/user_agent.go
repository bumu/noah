package uakit

import (
	"strings"
)

func IsValid(ua string) bool {

	if ua == "" ||
		strings.Contains(ua, "\\x") ||
		strings.Contains(ua, "bash") ||
		strings.Contains(ua, "echo") ||
		strings.Contains(ua, "JDatabaseDriverMysqli") ||
		strings.Contains(ua, "select") ||
		strings.Contains(ua, "SELECT") ||
		strings.Contains(ua, "union") ||
		strings.Contains(ua, "UNION") ||
		strings.Contains(ua, "CASE WHEN") ||
		strings.Contains(ua, "case when") ||
		strings.Contains(ua, "where") ||
		strings.Contains(ua, "WHERE") {

		return false
	}

	return true
}

func UseragetParse() {
}
