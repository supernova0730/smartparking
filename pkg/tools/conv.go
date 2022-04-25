package tools

import (
	"strconv"
	"strings"
	"time"
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

func GetDatePeriod(selectedPeriod string, dateFrom, dateTo time.Time) (from, to time.Time) {
	from, to = dateFrom, dateTo
	if from.IsZero() && to.IsZero() {
		from, to = GetSelectedDatePeriod(selectedPeriod, time.Now())
	}
	return StartOfDay(from), EndOfDay(to)
}

func GetSelectedDatePeriod(selectedPeriod string, dateTo time.Time) (from, to time.Time) {
	to = dateTo
	switch strings.ToLower(selectedPeriod) {
	case "all":
		from = time.Time{}
	case "month":
		from = to.AddDate(0, -1, 0)
	case "week":
		from = to.AddDate(0, 0, -7)
	case "today", "day":
		from = to.AddDate(0, 0, 0)
	default:
		from = to.AddDate(-1, 0, 0)
	}
	return
}

func StartOfDay(t time.Time) time.Time {
	y, m, d := t.Year(), t.Month(), t.Day()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

func EndOfDay(t time.Time) time.Time {
	y, m, d := t.Year(), t.Month(), t.Day()
	return time.Date(y, m, d, 23, 59, 59, 999999999, t.Location())
}
