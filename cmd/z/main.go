package main

import (
	"fmt"
	"os"
	"time"
)

const (
	timeFormatStr        = "2006-01-02 15:04:05"
	timeFormatStrUTC     = timeFormatStr + " UTC"
	timeFormatStrLocal   = timeFormatStr + " MST"
	timeFormatStrRFC3339 = timeFormatStr + "Z07:00"
)

const (
	timeSecond      = int64(time.Second)
	timeMillisecond = int64(time.Millisecond)
)

func main() {
	args := &binArgs{}
	msg, err := args.parse(nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse command line arguments: %v\n", err)
		os.Exit(1)
	}

	if len(msg) > 0 {
		fmt.Println(msg)
		os.Exit(0)
	}

	if len(args.Args.Timestamp) == 0 {
		if err := printNow(args); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		return
	}

	if err := parseTimestampAndPrintTime(args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func getTZ() (*time.Location, error) {
	tz := os.Getenv("TZ")
	if len(tz) > 0 {
		if tz == "UTC" || tz == "Local" {
			return time.UTC, nil
		}

		return time.LoadLocation(tz)
	}

	return time.Local, nil
}
