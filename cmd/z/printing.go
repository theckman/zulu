package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

type pData struct {
	timeUTC   time.Time
	timeLocal time.Time
	offset    float64
	timeSince time.Duration
}

func print(p pData) {
	unix := p.timeUTC.Unix()

	fmt.Printf("Local time:\t\t%s (%s)\n", p.timeLocal.Format(timeFormatStrLocal), alignmentStr(p.offset))

	fmt.Println()

	fmt.Printf("UTC time:\t\t%s", p.timeUTC.Format(timeFormatStrUTC))
	if p.timeSince >= time.Second || p.timeSince <= -time.Second {
		fmt.Printf(" [%s]", sinceStr(p.timeSince))
	}

	fmt.Println()

	fmt.Printf("Epoch timestamp:\t%d\n", unix)
	fmt.Printf("Millisecond timestamp:\t%d\n", secsToMilli(unix))
	fmt.Printf("Nanosecond timestamp:\t%d\n", unix*int64(time.Second))
}

func sinceStr(td time.Duration) string {
	ts := td / time.Second * time.Second

	if ts < 0 {
		ts *= -1
		return "occurs in " + ts.String()
	}

	return "was " + ts.String() + " ago"
}

func tzOffset(t time.Time) float64 {
	_, o := t.Zone()
	of := float64(o)

	if o == 0 {
		return of
	}

	// offset in seconds east of UTC to hours
	return of / 60.0 / 60.0
}

func secsToMilli(i int64) int64 {
	return (i * timeSecond) / timeMillisecond
}

func alignmentStr(offset float64) string {
	if offset == 0 {
		return "aligned with UTC"
	}

	offsetAbs := math.Abs(offset)

	unit := "hours"
	if offsetAbs == 1 {
		unit = "hour"
	}

	if offset > 0 {
		return fmt.Sprintf("%s %s ahead of UTC", strconv.FormatFloat(offset, 'f', -1, 64), unit)
	}

	return fmt.Sprintf("%s %s behind UTC", strconv.FormatFloat(offsetAbs, 'f', -1, 64), unit)
}

func printNow(args *binArgs) error {
	if args.Formatters.UnixEpoch {
		fmt.Printf("%d", time.Now().UTC().Unix())
		return nil
	}

	if args.Formatters.UnixMilli {
		fmt.Printf("%d", secsToMilli(time.Now().UTC().Unix()))
		return nil
	}

	if args.Formatters.UnixNano {
		fmt.Printf("%d", time.Now().UTC().Unix()*timeSecond)
		return nil
	}

	tz, err := getTZ()
	if err != nil {
		return fmt.Errorf("failed to load timezone %q: %v", tz, err)
	}

	tnu := time.Now().UTC()
	tnl := tnu.In(tz)

	if args.Formatters.RFC3339 {
		fmt.Print(tnu.Format(timeFormatStrRFC3339))
		return nil
	}

	if args.Formatters.ISO8601 {
		fmt.Print(tnu.Format(time.RFC3339))
		return nil
	}

	if args.Formatters.InternalTimeFormat {
		fmt.Print(tnu.Format(timeFormatStrLocal))
		return nil
	}

	if args.Formatters.Zulu {
		fmt.Print(tnu.Format(timeFormatStr))
		return nil
	}

	print(pData{
		timeUTC:   tnu,
		timeLocal: tnl,
		offset:    tzOffset(tnl),
	})

	return nil
}
