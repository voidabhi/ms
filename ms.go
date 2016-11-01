package ms

import (
  "errors"
	"regexp"
  "strings"
  "strconv"
  "fmt"
)

var s = float64(1000)
var m = float64(s) * 60
var h = float64(m) * 60
var d = float64(h) * 24
var y = float64(d) * 365.25

func Parse(str string) (float64, error) {
	re, err := regexp.Compile(`^((?:\d+)?\.?\d+) *(milliseconds?|msecs?|ms|seconds?|secs?|s|minutes?|mins?|m|hours?|hrs?|h|days?|d|years?|yrs?|y)?$`)
	if err != nil {
		return 0, err
	}

  matches := re.FindAllString(str, -1)
  if len(matches) == 0 {
    return 0, errors.New("Invalid input")
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
	case "years":
	case "year":
	case "yrs":
	case "yr":
	case "y":
		return n * y, nil
	case "days":
	case "day":
	case "d":
		return n * d, nil
	case "hours":
	case "hour":
	case "hrs":
	case "hr":
	case "h":
		return n * h, nil
	case "minutes":
	case "minute":
	case "mins":
	case "min":
	case "m":
		return n * m, nil
	case "seconds":
	case "second":
	case "secs":
	case "sec":
	case "s":
		return n * s, nil
	case "milliseconds":
	case "millisecond":
	case "msecs":
	case "msec":
	case "ms":
		return n, nil
	}
  return 0, nil
}
