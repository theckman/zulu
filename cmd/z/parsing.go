package main

import (
	"fmt"
	"strconv"
	"time"
)

func parseTimestampAndPrintTime(args *binArgs) error {
	tz, err := getTZ()
	if err != nil {
		return fmt.Errorf("failed to load timezone: %v", err)
	}

	tu, err := parseTimestamp(args)
	if err != nil {
		return fmt.Errorf("failed to parse provided timestamp: %v", err)
	}

	tul := tu.In(tz)

	print(pData{
		timeUTC:   tu,
		timeLocal: tul,
		offset:    tzOffset(tul),
		timeSince: time.Since(tul),
	})

	return nil
}

func parseTimestamp(args *binArgs) (time.Time, error) {
	s := args.Args.Timestamp

	switch {
	case args.Formatters.UnixEpoch, args.Formatters.UnixMilli, args.Formatters.UnixNano:
		return parseUnixTimestamp(s)
	case args.Formatters.RFC3339:
		return parseTime(timeFormatStrRFC3339, s)
	case args.Formatters.ISO8601:
		return parseTime(time.RFC3339, s)
	case args.Formatters.InternalTimeFormat:
		return parseTime(timeFormatStrLocal, s)
	case args.Formatters.Zulu:
		return time.ParseInLocation(timeFormatStr, s, time.UTC)
	default:
		return parseUnixTimestamp(s)
	}
}

func parseUnixTimestamp(s string) (time.Time, error) {
	ts, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, fmt.Errorf("provided timestamp (%q) does not appear to be a UNIX Timestamp", s)
	}

	strLen := len(s)

	switch {
	case strLen >= 19:
		ts = ts / timeSecond
	case strLen >= 13:
		ts = (ts * timeMillisecond) / timeSecond
	default:
		// i = i
	}

	return time.Unix(ts, 0).UTC(), nil
}

func parseTime(layout, value string) (time.Time, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return time.Time{}, err
	}

	return t.UTC(), nil
}
