package tools

import (
	"strconv"
)

func StringDefault(strVal, defaultVal string) string {
	if len(strVal) == 0 || strVal == "" {
		return defaultVal
	}
	return strVal
}

func StringToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}
