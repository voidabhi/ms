

package ms

import (
    "math"
    "regexp"
    "strconv"
)

var s float64 = 1000;
var m float64 = s * 60;
var h float64 = m * 60;
var d float64 = h * 24;
var y float64 = d * 365.25;

func Round(val float64, roundOn float64, places int ) (newVal float64) {
    var round float64
    pow := math.Pow(10, float64(places))
    digit := pow * val
    _, div := math.Modf(digit)
    if div >= roundOn {
        round = math.Ceil(digit)
    } else {
        round = math.Floor(digit)
    }
    newVal = round / pow
    return
}

func Parse(str string) {
  str = "" + str;
  if (len(str) > 10000) return;
  re, err := regexp.Compile(`(?i)^((?:\d+)?\.?\d+) *(milliseconds?|msecs?|ms|seconds?|secs?|s|minutes?|mins?|m|hours?|hrs?|h|days?|d|years?|yrs?|y)?$`)
  if err != nil {
      return ""
  }
  var n = parseFloat(match[1]);
  var type = (match[2] || 'ms').toLowerCase();
  switch (type) {
    case 'years':
    case 'year':
    case 'yrs':
    case 'yr':
    case 'y':
      return n * y;
    case 'days':
    case 'day':
    case 'd':
      return n * d;
    case 'hours':
    case 'hour':
    case 'hrs':
    case 'hr':
    case 'h':
      return n * h;
    case 'minutes':
    case 'minute':
    case 'mins':
    case 'min':
    case 'm':
      return n * m;
    case 'seconds':
    case 'second':
    case 'secs':
    case 'sec':
    case 's':
      return n * s;
    case 'milliseconds':
    case 'millisecond':
    case 'msecs':
    case 'msec':
    case 'ms':
      return n;
  }
}

func Short(ms float64) string {
  if (ms >= d) {
    return strconv.FormatFloat(Round(ms / d, .5, 0), 'f', -1, 64) + "d"
  }
  if (ms >= h) {
    return strconv.FormatFloat(Round(ms / h, .5, 0), 'f', -1, 64) + "h"
  }
  if (ms >= m) {
    return strconv.FormatFloat(Round(ms / m, .5, 0), 'f', -1, 64) + "m"
  }
  if (ms >= s) {
    return strconv.FormatFloat(Round(ms / s, .5, 0), 'f', -1, 64) + "s"
  }
  return strconv.FormatFloat(ms, 'f', -1, 64) + "ms"
}


func Long(ms float) string {
  return Plural(ms, d, 'day')
    || Plural(ms, h, 'hour')
    || Plural(ms, m, 'minute')
    || Plural(ms, s, 'second')
    || ms + ' ms';
}

/**
 * Pluralization helper.
 */

func Plural(ms float64, n float64, name string) (string) {
  if (ms < n) {
    return ""
  }

  if (ms < n * 1.5) {
    return  strconv.FormatFloat(math.Floor(ms / n), 'f', -1, 64) + " " + name
  }
  return strconv.FormatFloat(math.Ceil(ms / n), 'f', -1, 64) + " " + name + "s"
}
