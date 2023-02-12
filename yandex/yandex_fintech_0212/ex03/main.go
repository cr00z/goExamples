package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	WEEK    = "WEEK"
	MONTH   = "MONTH"
	QUARTER = "QUARTER"
	YEAR    = "YEAR"
	REVIEW  = "REVIEW"
)

func getInput() (time.Time, time.Time, string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	intervalType := scanner.Text()

	scanner.Scan()
	interval := strings.Split(scanner.Text(), " ")
	start, _ := time.Parse("2006-01-02", interval[0])
	end, _ := time.Parse("2006-01-02", interval[1])

	return start, end, intervalType
}

func getNextEnd(start time.Time, intervalType string) time.Time {
	next := start
	month := start.Month()
	year := start.Year()

	switch intervalType {

	case WEEK:
		for next.Weekday() != time.Sunday {
			next = next.AddDate(0, 0, 1)
		}

	case MONTH:
		if month == time.December {
			month = time.January
			year++
		} else {
			month++
		}
		next = time.Date(year, month, 0, 0, 0, 0, 0, time.UTC)

	case QUARTER:
		switch month {
		case time.January, time.February, time.March:
			month = time.April
		case time.April, time.May, time.June:
			month = time.July
		case time.July, time.August, time.September:
			month = time.October
		case time.October, time.November, time.December:
			month = time.January
			year++
		}
		next = time.Date(year, month, 0, 0, 0, 0, 0, time.UTC)

	case YEAR:
		next = time.Date(year+1, 1, 0, 0, 0, 0, 0, time.UTC)

	case REVIEW:
		if month >= time.April && month <= time.September {
			month = time.October
		} else {
			if month > time.September {
				year++
			}
			month = time.April
		}
		next = time.Date(year, month, 0, 0, 0, 0, 0, time.UTC)
	}

	return next
}

func outInterval(out *strings.Builder, start time.Time, end time.Time) {
	out.WriteString(start.Format("2006-01-02"))
	out.WriteByte(' ')
	out.WriteString(end.Format("2006-01-02"))
	out.WriteByte('\n')
}

func fixFebruary(dt time.Time) time.Time {
	for dt.Day() != 1 {
		dt = dt.AddDate(0, 0, 1)
	}
	return dt.AddDate(0, 0, -1)
}

func main() {
	start, end, intervalType := getInput()

	count := 0
	var out strings.Builder

	nextEnd := getNextEnd(start, intervalType)
	for nextEnd.Before(end) || nextEnd.Equal(end) {
		count++
		outInterval(&out, start, nextEnd)

		start = nextEnd.AddDate(0, 0, 1)

		switch intervalType {
		case WEEK:
			nextEnd = nextEnd.AddDate(0, 0, 7)

		case MONTH:
			nextEnd = fixFebruary(nextEnd.AddDate(0, 0, 28))

		case QUARTER:
			nextEnd = fixFebruary(nextEnd.AddDate(0, 2, 28))

		case YEAR:
			nextEnd = nextEnd.AddDate(1, 0, 0)

		case REVIEW:
			nextEnd = fixFebruary(nextEnd.AddDate(0, 5, 28))
		}
	}

	if start.Before(end) || start.Equal(end) {
		count++
		outInterval(&out, start, end)
	}

	fmt.Println(count)
	fmt.Println(out.String())
}
