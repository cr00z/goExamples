package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FormatDuration(seconds int64) string {
	var result string
	periods := []string{" year", " day", " hour", " minute", " second"}
	var times [5]int64
	times[0] = seconds / (365 * 24 * 3600)
	seconds -= times[0] * (365 * 24 * 3600)
	times[1] = seconds / (24 * 3600)
	seconds -= times[1] * (24 * 3600)
	times[2] = seconds / 3600
	seconds -= times[2] * 3600
	times[3] = seconds / 60
	times[4] = seconds - times[3]*60
	var notZero int
	for i := 0; i < 5; i++ {
		if times[i] > 0 {
			notZero += 1
			result += strconv.Itoa(int(times[i])) + periods[i]
			if times[i] > 1 {
				result += "s"
			}
			result += ", "
		}
	}
	result = strings.TrimRight(result, ", ")
	if notZero > 1 {
		i := strings.LastIndex(result, ", ")
		result = result[:i] + strings.Replace(result[i:], ", ", " and ", 1)
	}
	return result
}

func main() {
	fmt.Println(FormatDuration(3662))
}
