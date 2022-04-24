package tools

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	dayInHours   = 24
	weekInHours  = 168
	monthInHours = 24 * 30
	yearInHours  = 24 * 30 * 12
)

var unitMap = map[string]int64{
	"h": int64(time.Hour),                // hour
	"d": int64(time.Hour) * dayInHours,   // day
	"w": int64(time.Hour) * weekInHours,  // week
	"m": int64(time.Hour) * monthInHours, // month
	"y": int64(time.Hour) * yearInHours,  // year
}

func DurationToString(d time.Duration) string {
	hours := int(d.Seconds() / 3600)
	days := hours / dayInHours
	hours %= dayInHours
	years := days / (yearInHours / dayInHours)
	days %= yearInHours / dayInHours
	months := days / (monthInHours / dayInHours)
	days %= monthInHours / dayInHours
	weeks := days / (weekInHours / dayInHours)
	days %= weekInHours / dayInHours

	var result strings.Builder
	if years > 0 {
		result.WriteString(fmt.Sprintf("%dy", years))
	}
	if months > 0 {
		result.WriteString(fmt.Sprintf("%dm", months))
	}
	if weeks > 0 {
		result.WriteString(fmt.Sprintf("%dw", weeks))
	}
	if days > 0 {
		result.WriteString(fmt.Sprintf("%dd", days))
	}
	if hours > 0 {
		result.WriteString(fmt.Sprintf("%dh", hours))
	}

	return result.String()
}

func ParseDuration(s string) (time.Duration, error) {
	if len(s) == 0 {
		return 0, nil
	}

	orig := s
	var d int64
	neg := false

	if s != "" {
		c := s[0]
		if c == '-' || c == '+' {
			neg = c == '-'
			s = s[1:]
		}
	}
	if s == "0" {
		return 0, nil
	}
	if s == "" {
		return 0, fmt.Errorf("time: invalid duration %q", orig)
	}
	for s != "" {
		var (
			v, f  int64
			scale float64 = 1
		)

		var err error

		if !(s[0] == '.' || '0' <= s[0] && s[0] <= '9') {
			return 0, fmt.Errorf("time: invalid duration %q", orig)
		}
		pl := len(s)
		v, s, err = leadingInt(s)
		if err != nil {
			return 0, fmt.Errorf("time: invalid duration %q", orig)
		}
		pre := pl != len(s)

		post := false
		if s != "" && s[0] == '.' {
			s = s[1:]
			pl := len(s)
			f, scale, s = leadingFraction(s)
			post = pl != len(s)
		}
		if !pre && !post {
			return 0, fmt.Errorf("time: invalid duration %q", orig)
		}

		i := 0
		for ; i < len(s); i++ {
			c := s[i]
			if c == '.' || '0' <= c && c <= '9' {
				break
			}
		}
		if i == 0 {
			return 0, fmt.Errorf("time: missing unit in duration %q", orig)
		}
		u := s[:i]
		s = s[i:]
		unit, ok := unitMap[u]
		if !ok {
			return 0, fmt.Errorf("time: unknown unit %q in duration %q", u, orig)
		}
		if v > (1<<63-1)/unit {
			return 0, fmt.Errorf("time: invalid duration %q", orig)
		}
		v *= unit
		if f > 0 {
			v += int64(float64(f) * (float64(unit) / scale))
			if v < 0 {
				return 0, fmt.Errorf("time: invalid duration %q", orig)
			}
		}
		d += v
		if d < 0 {
			return 0, fmt.Errorf("time: invalid duration %q", orig)
		}
	}

	if neg {
		d = -d
	}
	return time.Duration(d), nil
}

var errLeadingInt = errors.New("time: bad [0-9]*")

func leadingInt(s string) (x int64, rem string, err error) {
	i := 0
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			break
		}
		if x > (1<<63-1)/10 {
			return 0, "", errLeadingInt
		}
		x = x*10 + int64(c) - '0'
		if x < 0 {
			return 0, "", errLeadingInt
		}
	}
	return x, s[i:], nil
}

func leadingFraction(s string) (x int64, scale float64, rem string) {
	i := 0
	scale = 1
	overflow := false
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			break
		}
		if overflow {
			continue
		}
		if x > (1<<63-1)/10 {
			overflow = true
			continue
		}
		y := x*10 + int64(c) - '0'
		if y < 0 {
			overflow = true
			continue
		}
		x = y
		scale *= 10
	}
	return x, scale, s[i:]
}
