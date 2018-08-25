package main

import (
	"fmt"
	"os"
	"runtime"

	flags "github.com/jessevdk/go-flags"
)

const appVersion = "0.1.0"

const appVersionString = `zulu v%s built with %s
Copyright 2018 Tim Heckman
Released under the MIT License`

type binArgs struct {
	Version bool `short:"V" long:"version" description:"print version and exit"`

	Formatters struct {
		UnixEpoch          bool `short:"u" long:"unix" description:"time format is the UNIX Epoch (seconds since 00:00:00 UTC January 1, 1970)"`
		UnixMilli          bool `short:"m" long:"unix-milliseconds" description:"time format is the UNIX Epoch to the nearest second with millisecond precision"`
		UnixNano           bool `short:"n" long:"unix-nanoseconds" description:"time format is the UNIX Epoch to the nearest second with nanosecond precision"`
		RFC3339            bool `short:"r" long:"rfc-3339" description:"time format is the RFC 3339 format in UTC timezone, with the T replaced by a space"`
		ISO8601            bool `short:"i" long:"iso-8601" description:"time format is the same as the -r/--rfc-3339 format, except it includes the 'T' between the date and time"`
		InternalTimeFormat bool `short:"t" long:"internal-time-format" description:"time format is the same as is printed by default (e.g. 2018-08-27 15:00:00 UTC)"`
		Zulu               bool `short:"z" long:"zulu" description:"same as -r/--rfc-3339 but with no timezone specifier"`
	} `group:"Timestamp Formatters"`

	// Args are the positional arguments provided on the command line. By
	// default it accepts a timestamp whose format depends on which Formatter
	// flags you use.
	Args struct {
		Timestamp string `positional-arg-name:"TIMESTAMP" description:"An optional timestamp to parse and print the datetime of. The expected format depends on which Timestamp Formatter flag you use. The default is a UNIX Epoch parser that attempts to assume the desired precision (seconds, milliseconds, nanoseconds)"`
	} `positional-args:"yes"`
}

func (a *binArgs) parse(args []string) (string, error) {
	if args == nil {
		args = os.Args
	}

	p := flags.NewParser(a, flags.HelpFlag|flags.PassDoubleDash)

	_, err := p.ParseArgs(args[1:])

	// determine if there was a parsing error
	// unfortunately, help message is returned as an error
	if err != nil {
		// determine whether this was a help message by doing a type
		// assertion of err to *flags.Error and check the error type
		// if it was a help message, do not return an error
		if errType, ok := err.(*flags.Error); ok {
			if errType.Type == flags.ErrHelp {
				return err.Error(), nil
			}
		}

		return "", err
	}

	if a.Version {
		out := fmt.Sprintf(
			appVersionString,
			appVersion, runtime.Version(),
		)
		return out, nil
	}

	return "", nil
}
