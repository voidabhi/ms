package ms

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var s = float64(1000)
var m = float64(s) * 60
var h = float64(m) * 60
var d = float64(h) * 24
var y = float64(d) * 365.25

var (
	errInvalidInput = errors.New("Invalid Input")
)

// Parse : converts time formats to milliseconds
func Parse(time string) (float64, error) {
	re, err := regexp.Compile(`((?:\d+)?\.?\d+) *(milliseconds?|msecs?|ms|seconds?|secs?|s|minutes?|mins?|m|hours?|hrs?|h|days?|d|years?|yrs?|y)?`)
	if err != nil {
		return 0, err
	}

	var matches []string
	for i, match := range re.FindStringSubmatch(time) {
		if i > 0 {
			matches = append(matches, match)
		}
	}

	n, err := strconv.ParseFloat(matches[0], 64)
	if err != nil {
		return 0, err
	}

	if len(matches) == 1 {
		matches = append(matches, "ms")
	}

	t := strings.ToLower(matches[1])

	switch t {
	case "years", "year", "yrs", "yr", "y":
		return n * y, nil
	case "days", "day", "d":
		return n * d, nil
	case "hours", "hour", "hrs", "hr", "h":
		return n * h, nil
	case "minutes", "minute", "mins", "min", "m":
		return n * m, nil
	case "seconds", "second", "secs", "sec", "s":
		return n * s, nil
	case "milliseconds", "millisecond", "msecs", "msec", "ms":
		return n, nil
	}

	return 0, errInvalidInput
}
